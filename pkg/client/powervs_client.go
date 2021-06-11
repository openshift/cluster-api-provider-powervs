package client

import (
	"context"
	"fmt"
	gohttp "net/http"
	"strings"
	"time"

	"github.com/IBM-Cloud/bluemix-go"
	"github.com/IBM-Cloud/bluemix-go/api/resource/resourcev2/controllerv2"
	"github.com/IBM-Cloud/bluemix-go/authentication"
	"github.com/IBM-Cloud/bluemix-go/http"
	bluemixmodels "github.com/IBM-Cloud/bluemix-go/models"
	"github.com/IBM-Cloud/bluemix-go/rest"
	bxsession "github.com/IBM-Cloud/bluemix-go/session"
	"github.com/IBM-Cloud/power-go-client/clients/instance"
	"github.com/IBM-Cloud/power-go-client/ibmpisession"
	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/models"
	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"

	"github.com/dgrijalva/jwt-go"
	machineapiapierrors "github.com/openshift/machine-api-operator/pkg/controller/machine"
	"github.com/pkg/errors"
	utils "github.com/ppc64le-cloud/powervs-utils"
	corev1 "k8s.io/api/core/v1"
	apimachineryerrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	TIMEOUT = time.Hour

	DefaultCredentialNamespace = "openshift-machine-api"
	DefaultCredentialSecret    = "powervs-credentials-secret"

	InstanceStateNameShutoff = "SHUTOFF"
	InstanceStateNameActive  = "ACTIVE"
	InstanceStateNameBuild   = "BUILD"

	PowerServiceType = "power-iaas"
)

var (
	ErrorInstanceNotFound = errors.New("Instance Not Found")
)

func FormatProviderID(instanceID string) string {
	return fmt.Sprintf("powervs:///%s", instanceID)
}

type PowerVSClientBuilderFuncType func(client client.Client, secretName, namespace, cloudInstanceID, region string) (Client, error)

func apiKeyFromSecret(secret *corev1.Secret) (apiKey string, err error) {
	switch {
	case len(secret.Data["IBMCLOUD_API_KEY"]) > 0:
		apiKey = string(secret.Data["IBMCLOUD_API_KEY"])
	default:
		return "", fmt.Errorf("invalid secret for powervs credentials")
	}
	return
}

func GetAPIKey(ctrlRuntimeClient client.Client, secretName, namespace string) (apikey string, err error) {
	if secretName == "" {
		return "", machineapiapierrors.InvalidMachineConfiguration("empty secret name")
	}
	var secret corev1.Secret
	if err := ctrlRuntimeClient.Get(context.Background(), client.ObjectKey{Namespace: namespace, Name: secretName}, &secret); err != nil {
		if apimachineryerrors.IsNotFound(err) {
			return "", machineapiapierrors.InvalidMachineConfiguration("powervs credentials secret %s/%s: %v not found", namespace, secretName, err)
		}
		return "", err
	}
	apikey, err = apiKeyFromSecret(&secret)
	if err != nil {
		return "", fmt.Errorf("failed to create shared credentials file from Secret: %v", err)
	}
	return
}

// getServiceURL returns the appropriate service URL for the VPC for given region or error
func getServiceURL(region string) (string, error) {
	switch region {
	case "us-south", "us-east", "eu-gb", "eu-de", "au-syd", "jp-tok", "jp-osa", "ca-tor":
		return fmt.Sprintf("https://%s.iaas.cloud.ibm.com/v1", region), nil
	default:
		return "", fmt.Errorf("invalid region: %s", region)
	}
}

func NewValidatedClient(ctrlRuntimeClient client.Client, secretName, namespace, cloudInstanceID, region string) (Client, error) {
	apikey, err := GetAPIKey(ctrlRuntimeClient, secretName, namespace)
	if err != nil {
		return nil, err
	}

	s, err := bxsession.New(&bluemix.Config{BluemixAPIKey: apikey})
	if err != nil {
		return nil, err
	}

	// Set the region as us-south if not set
	if region == "" {
		region = "us-south"
	}

	url, err := getServiceURL(region)
	if err != nil {
		return nil, err
	}

	vpcClient, err := vpcv1.NewVpcV1(&vpcv1.VpcV1Options{
		Authenticator: &core.IamAuthenticator{
			ApiKey: apikey,
		},
		URL: url,
	})

	if err != nil {
		return nil, err
	}

	c := &powerVSClient{
		cloudInstanceID: cloudInstanceID,
		Session:         s,
		VPCClient:       vpcClient,
	}

	err = authenticateAPIKey(s)
	if err != nil {
		return c, err
	}

	c.User, err = fetchUserDetails(s, 2)
	if err != nil {
		return c, err
	}

	ctrlv2, err := controllerv2.New(s)
	if err != nil {
		return c, err
	}

	c.ResourceClient = ctrlv2.ResourceServiceInstanceV2()

	resource, err := c.ResourceClient.GetInstance(cloudInstanceID)
	if err != nil {
		return nil, err
	}
	r, err := utils.GetRegion(resource.RegionID)
	if err != nil {
		return nil, err
	}
	zone := resource.RegionID

	c.session, err = ibmpisession.New(c.Config.IAMAccessToken, r, true, time.Hour, c.User.Account, zone)
	if err != nil {
		return nil, err
	}

	c.InstanceClient = instance.NewIBMPIInstanceClient(c.session, cloudInstanceID)
	c.NetworkClient = instance.NewIBMPINetworkClient(c.session, cloudInstanceID)
	return c, err
}

// NewClientMinimal is bare minimal client can be used for quarrying the resources
func NewClientMinimal(apiKey string) (Client, error) {
	s, err := bxsession.New(&bluemix.Config{BluemixAPIKey: apiKey})
	if err != nil {
		return nil, err
	}

	c := &powerVSClient{
		Session: s,
	}

	ctrlv2, err := controllerv2.New(s)
	if err != nil {
		return c, err
	}

	c.ResourceClient = ctrlv2.ResourceServiceInstanceV2()

	return c, nil
}

type powerVSClient struct {
	region          string
	zone            string
	cloudInstanceID string

	*bxsession.Session
	VPCClient      *vpcv1.VpcV1
	User           *User
	ResourceClient controllerv2.ResourceServiceInstanceRepository
	session        *ibmpisession.IBMPISession
	InstanceClient *instance.IBMPIInstanceClient
	NetworkClient  *instance.IBMPINetworkClient
}

func (p *powerVSClient) DeleteInstance(id string) error {
	return p.InstanceClient.Delete(id, p.cloudInstanceID, TIMEOUT)
}

func (p *powerVSClient) CreateInstance(params *p_cloud_p_vm_instances.PcloudPvminstancesPostParams) (*models.PVMInstanceList, error) {
	return p.InstanceClient.Create(params, p.cloudInstanceID, TIMEOUT)
}

func (p *powerVSClient) GetInstance(id string) (*models.PVMInstance, error) {
	return p.InstanceClient.Get(id, p.cloudInstanceID, TIMEOUT)
}

func (p *powerVSClient) GetInstanceByName(name string) (*models.PVMInstance, error) {
	instances, err := p.GetInstances()
	if err != nil {
		return nil, fmt.Errorf("failed to get the instance list")
	}

	for _, i := range instances.PvmInstances {
		if *i.ServerName == name {
			return p.GetInstance(*i.PvmInstanceID)
		}
	}
	return nil, ErrorInstanceNotFound
}

func (p *powerVSClient) GetInstances() (*models.PVMInstances, error) {
	return p.InstanceClient.GetAll(p.cloudInstanceID, TIMEOUT)
}

func (p *powerVSClient) GetCloudServiceInstances() ([]bluemixmodels.ServiceInstanceV2, error) {
	var instances []bluemixmodels.ServiceInstanceV2
	svcs, err := p.ResourceClient.ListInstances(controllerv2.ServiceInstanceQuery{
		Type: "service_instance",
	})
	if err != nil {
		return svcs, fmt.Errorf("failed to list the service instances: %v", err)
	}
	for _, svc := range svcs {
		if svc.Crn.ServiceName == PowerServiceType {
			instances = append(instances, svc)
		}
	}
	return instances, nil
}

func authenticateAPIKey(sess *bxsession.Session) error {
	config := sess.Config
	tokenRefresher, err := authentication.NewIAMAuthRepository(config, &rest.Client{
		DefaultHeader: gohttp.Header{
			"User-Agent": []string{http.UserAgent()},
		},
	})
	if err != nil {
		return err
	}
	return tokenRefresher.AuthenticateAPIKey(config.BluemixAPIKey)
}

type User struct {
	ID         string
	Email      string
	Account    string
	cloudName  string `default:"bluemix"`
	cloudType  string `default:"public"`
	generation int    `default:"2"`
}

func fetchUserDetails(sess *bxsession.Session, generation int) (*User, error) {
	config := sess.Config
	user := User{}
	var bluemixToken string

	if strings.HasPrefix(config.IAMAccessToken, "Bearer") {
		bluemixToken = config.IAMAccessToken[7:len(config.IAMAccessToken)]
	} else {
		bluemixToken = config.IAMAccessToken
	}

	token, err := jwt.Parse(bluemixToken, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})
	if err != nil && !strings.Contains(err.Error(), "key is of invalid type") {
		return &user, err
	}

	claims := token.Claims.(jwt.MapClaims)
	if email, ok := claims["email"]; ok {
		user.Email = email.(string)
	}
	user.ID = claims["id"].(string)
	user.Account = claims["account"].(map[string]interface{})["bss"].(string)
	iss := claims["iss"].(string)
	if strings.Contains(iss, "https://iam.cloud.ibm.com") {
		user.cloudName = "bluemix"
	} else {
		user.cloudName = "staging"
	}
	user.cloudType = "public"

	user.generation = generation
	return &user, nil
}
