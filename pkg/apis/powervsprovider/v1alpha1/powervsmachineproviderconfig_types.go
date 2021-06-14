/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PowerVSMachineProviderConfig is the Schema for the powervsmachineproviderconfigs API
type PowerVSMachineProviderConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Region for VPC client for load balancer
	Region string `json:"region,omitempty"`

	// LoadBalancers is the set of load balancers to which the new instance
	// should be added once it is created.
	LoadBalancers []LoadBalancerReference `json:"loadBalancers,omitempty"`

	// ServiceInstanceID is the PowerVS service ID
	ServiceInstanceID string `json:"serviceInstanceID"`

	// ImageID is the Image ID used for deploying the machine
	ImageID string `json:"imageID"`

	// UserDataSecret is the k8s secret contains the user data script
	UserDataSecret *corev1.LocalObjectReference `json:"userDataSecret,omitempty"`

	// CredentialsSecret is the k8s secret contains the API Key for IBM Cloud authentication
	CredentialsSecret *corev1.LocalObjectReference `json:"credentialsSecret,omitempty"`

	// MachineType is the System type used to host the vsi
	MachineType string `json:"machineType"`

	// ProcessorType is the processor type, e.g: dedicated, shared, capped
	ProcessorType string `json:"processorType"`

	// Cores is Number of processors allocated
	Cores string `json:"cores"`

	// Memory is Amount of memory allocated (in GB)
	Memory string `json:"memory"`

	// Subnets is an array of network to be attached to the machine
	Subnets []string `json:"subnets"`

	// KeyName is the name of the KeyPair to use for SSH
	KeyName *string `json:"keyName,omitempty"`
}

// LoadBalancerReference is a reference to a load balancer on PowerVS.
type LoadBalancerReference struct {
	// The load balancer identifier.
	ID string `json:"id"`

	// The pool identifier.
	PoolID string `json:"poolID"`

	// The port number of the application running in the server member(machine)
	Port int64 `json:"port"`
}

//+kubebuilder:object:root=true

// PowerVSMachineProviderConfigList contains a list of PowerVSMachineProviderConfig
type PowerVSMachineProviderConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PowerVSMachineProviderConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PowerVSMachineProviderConfig{}, &PowerVSMachineProviderConfigList{}, &PowerVSMachineProviderStatus{})
}
