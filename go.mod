module github.com/openshift/cluster-api-provider-powervs

go 1.15

require (
	github.com/IBM-Cloud/bluemix-go v0.0.0-20211102075456-ffc4e11dfb16
	github.com/IBM-Cloud/power-go-client v1.0.76
	github.com/IBM/go-sdk-core/v5 v5.8.0
	github.com/blang/semver v3.5.1+incompatible
	github.com/golang-jwt/jwt/v4 v4.1.0
	github.com/golang/mock v1.5.0
	github.com/onsi/gomega v1.16.0
	github.com/openshift/api v0.0.0-20211028135425-c4970133b5ba
	github.com/openshift/machine-api-operator v0.2.1-0.20211029132328-128c5c90918c
	github.com/pkg/errors v0.9.1
	github.com/ppc64le-cloud/powervs-utils v0.0.0-20210415051532-4cdd6a79c8fa

	// kube 1.22
	k8s.io/api v0.22.1
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
	k8s.io/klog/v2 v2.9.0
	sigs.k8s.io/controller-runtime v0.9.6
	sigs.k8s.io/controller-tools v0.6.3-0.20210916130746-94401651a6c3
	sigs.k8s.io/yaml v1.2.0
)
