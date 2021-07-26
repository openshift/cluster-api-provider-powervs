package machine

import (
	"fmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
	"k8s.io/apimachinery/pkg/util/rand"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openshift/cluster-api-provider-powervs/pkg/apis/powervsprovider/v1alpha1"
	machinev1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
)

const (
	defaultNamespace      = "default"
	credentialsSecretName = "powervs-credentials-secret"
	userDataSecretName    = "powervs-actuator-user-data-secret"
	nameLength            = 5
)

func stubUserDataSecret(name string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: defaultNamespace,
		},
		Data: map[string][]byte{
			userDataSecretKey: []byte("userDataBlob"),
		},
	}
}

func stubPowerVSCredentialsSecret(name string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: defaultNamespace,
		},
		Data: map[string][]byte{
			"ibmcloud_api_key": []byte("Kl9k1elFgPb_QgEDF0d5iNHMOFa--YX6JWLpi0XkWn"),
		},
	}
}

func stubMachine() (*machinev1.Machine, error) {

	credSecretName := fmt.Sprintf("%s-%s", credentialsSecretName, rand.String(nameLength))
	providerSpec, err := v1alpha1.RawExtensionFromProviderSpec(stubProviderConfig(credSecretName))
	if err != nil {
		return nil, err
	}

	return &machinev1.Machine{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: defaultNamespace,
			Labels: map[string]string{
				machinev1.MachineClusterIDLabel: "CLUSTERID",
			},
		},
		Spec: machinev1.MachineSpec{
			ProviderSpec: machinev1.ProviderSpec{
				Value: providerSpec,
			},
		}}, nil
}

func stubProviderConfig(name string) *v1alpha1.PowerVSMachineProviderConfig {
	testKeyPair := "Test-KeyPair"
	return &v1alpha1.PowerVSMachineProviderConfig{
		CredentialsSecret: &corev1.LocalObjectReference{
			Name: name,
		},
		Memory:      "32",
		Processors:  "0.5",
		KeyPairName: &testKeyPair,
	}
}

func stubGetInstances() *models.PVMInstanceList {
	return &models.PVMInstanceList{stubGetInstance()}
}

func stubGetInstance() *models.PVMInstance {
	dummyInstanceID := "instance-id"
	status := "ACTIVE"
	return &models.PVMInstance{
		PvmInstanceID: &dummyInstanceID,
		Status:        &status,
	}
}
