package machine

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/IBM-Cloud/power-go-client/power/client/p_cloud_p_vm_instances"
	"github.com/IBM-Cloud/power-go-client/power/models"
	powervsproviderv1 "github.com/openshift/cluster-api-provider-powervs/pkg/apis/powervsprovider/v1alpha1"
	powervsclient "github.com/openshift/cluster-api-provider-powervs/pkg/client"
	machinev1 "github.com/openshift/machine-api-operator/pkg/apis/machine/v1beta1"
	mapierrors "github.com/openshift/machine-api-operator/pkg/controller/machine"
	"k8s.io/klog/v2"
)

// removeStoppedMachine removes all instances of a specific machine that are in a stopped state.
func removeStoppedMachine(machine *machinev1.Machine, client powervsclient.Client) error {
	instance, err := client.GetInstanceByName(machine.Name)
	if err != nil && err != powervsclient.ErrorInstanceNotFound {
		klog.Errorf("Error getting instance by name: %s, err: %v", machine.Name, err)
		return fmt.Errorf("error getting instance by name: %s, err: %v", machine.Name, err)
	} else if err == powervsclient.ErrorInstanceNotFound {
		klog.Infof("Instance not found with name: %n", machine.Name)
		return nil
	}

	if instance != nil && *instance.Status == powervsclient.InstanceStateNameShutoff {
		return client.DeleteInstance(*instance.PvmInstanceID)
	}
	return nil
}

func launchInstance(machine *machinev1.Machine, machineProviderConfig *powervsproviderv1.PowerVSMachineProviderConfig, userData []byte, client powervsclient.Client) (*models.PVMInstance, error) {
	// code for powervs

	memory, err := strconv.ParseFloat(machineProviderConfig.Memory, 64)
	if err != nil {
		return nil, mapierrors.InvalidMachineConfiguration("failed to convert memory(%s) to float64", machineProviderConfig.Memory)
	}
	cores, err := strconv.ParseFloat(machineProviderConfig.Cores, 64)
	if err != nil {
		return nil, mapierrors.InvalidMachineConfiguration("failed to convert Cores(%s) to float64", machineProviderConfig.Cores)
	}

	var nets []*models.PVMInstanceAddNetwork

	for _, net := range machineProviderConfig.Subnets {
		nets = append(nets, &models.PVMInstanceAddNetwork{NetworkID: &net})
	}

	params := &p_cloud_p_vm_instances.PcloudPvminstancesPostParams{
		Body: &models.PVMInstanceCreate{
			ImageID:     &machineProviderConfig.ImageID,
			KeyPairName: *machineProviderConfig.KeyName,
			Networks:    nets,
			ServerName:  &machine.Name,
			Memory:      &memory,
			Processors:  &cores,
			ProcType:    &machineProviderConfig.ProcessorType,
			SysType:     machineProviderConfig.MachineType,
			UserData:    base64.StdEncoding.EncodeToString(userData),
		},
	}
	instances, err := client.CreateInstance(params)
	if err != nil {
		return nil, mapierrors.CreateMachine("error creating powervs instance: %v", err)
	}

	insIDs := make([]string, 0)
	for _, in := range *instances {
		insID := in.PvmInstanceID
		insIDs = append(insIDs, *insID)
	}

	if len(insIDs) == 0 {
		return nil, mapierrors.CreateMachine("error getting the instance ID post deployment for: %s", machine.Name)
	}

	instance, err := client.GetInstance(insIDs[0])
	if err != nil {
		return nil, mapierrors.CreateMachine("error getting the instance for ID: %s", insIDs[0])
	}
	return instance, nil
}
