module github.com/openshift/cluster-api-provider-powervs

go 1.15

require (
	github.com/IBM-Cloud/bluemix-go v0.0.0-20200921095234-26d1d0148c62
	github.com/IBM-Cloud/power-go-client v1.0.55
	github.com/IBM/go-sdk-core/v4 v4.5.1
	github.com/IBM/vpc-go-sdk v1.0.1
	github.com/aws/aws-sdk-go v1.35.20
	github.com/blang/semver v3.5.1+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/mock v1.4.4
	github.com/onsi/ginkgo v1.15.0 // indirect
	github.com/onsi/gomega v1.10.5 // indirect
	github.com/openshift/api v0.0.0-20201216151826-78a19e96f9eb
	github.com/openshift/machine-api-operator v0.2.1-0.20210104142355-8e6ae0acdfcf
	github.com/pkg/errors v0.9.1
	github.com/ppc64le-cloud/powervs-utils v0.0.0-20210415051532-4cdd6a79c8fa
	golang.org/x/tools v0.1.0 // indirect

	// kube 1.18
	k8s.io/api v0.20.0
	k8s.io/apimachinery v0.20.0
	k8s.io/client-go v0.20.0
	k8s.io/klog/v2 v2.4.0
	sigs.k8s.io/controller-runtime v0.7.0
	sigs.k8s.io/controller-tools v0.3.0
	sigs.k8s.io/yaml v1.2.0
)

replace (
	sigs.k8s.io/cluster-api-provider-aws => github.com/openshift/cluster-api-provider-aws v0.2.1-0.20201216171336-0b00fb8d96ac
	sigs.k8s.io/cluster-api-provider-azure => github.com/openshift/cluster-api-provider-azure v0.1.0-alpha.3.0.20201209184807-075372e2ed03
)
