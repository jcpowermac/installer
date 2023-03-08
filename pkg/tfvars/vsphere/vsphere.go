package vsphere

import (
	"encoding/json"
	"fmt"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/vim25"
	"golang.org/x/net/context"
	"time"

	"github.com/pkg/errors"

	machineapi "github.com/openshift/api/machine/v1beta1"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/installconfig/vsphere"
	"github.com/openshift/installer/pkg/tfvars/internal/cache"
	vtypes "github.com/openshift/installer/pkg/types/vsphere"
)

type folder struct {
	Name       string `json:"name"`
	Datacenter string `json:"vsphere_datacenter"`
}

type config struct {
	OvaFilePath              string                                   `json:"vsphere_ova_filepath"`
	DiskType                 vtypes.DiskType                          `json:"vsphere_disk_type"`
	VCenters                 map[string]vtypes.VCenter                `json:"vsphere_vcenters"`
	FailureDomains           []vtypes.FailureDomain                   `json:"vsphere_failure_domains"`
	NetworksInFailureDomains map[string]string                        `json:"vsphere_networks"`
	ControlPlanes            []*machineapi.VSphereMachineProviderSpec `json:"vsphere_control_planes"`
	DatacentersFolders       map[string]*folder                       `json:"vsphere_folders"`
}

// TFVarsSources contains the parameters to be converted into Terraform variables
type TFVarsSources struct {
	ControlPlaneConfigs     []*machineapi.VSphereMachineProviderSpec
	ImageURL                string
	DiskType                vtypes.DiskType
	NetworksInFailureDomain map[string]string
	InstallConfig           *installconfig.InstallConfig
	InfraID                 string
	ControlPlaneMachines    []machineapi.Machine
}

// TFVars generate vSphere-specific Terraform variables
func TFVars(sources TFVarsSources) ([]byte, error) {
	cachedImage, err := cache.DownloadImageFile(sources.ImageURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to use cached vsphere image")
	}

	vcenterZones := convertVCentersToMap(sources.InstallConfig.Config.VSphere.VCenters)
	datacentersFolders, err := createDatacenterFolderMap(sources.InfraID, sources.InstallConfig.Config.VSphere.FailureDomains, vcenterZones)
	if err != nil {
		return nil, err
	}

	cfg := &config{
		OvaFilePath:              cachedImage,
		DiskType:                 sources.DiskType,
		VCenters:                 vcenterZones,
		FailureDomains:           sources.InstallConfig.Config.VSphere.FailureDomains,
		NetworksInFailureDomains: sources.NetworksInFailureDomain,
		ControlPlanes:            sources.ControlPlaneConfigs,
		DatacentersFolders:       datacentersFolders,
	}

	return json.MarshalIndent(cfg, "", "  ")
}

// createDatacenterFolderMap()
// This function loops over the range of failure domains
// Each failure domain defines the vCenter datacenter and folder
// to be used for the virtual machines within that domain.
// The datacenter could be reused but a folder could be
// unique - the key then becomes a string that contains
// both the datacenter name and the folder to be created.

func createDatacenterFolderMap(infraID string, failureDomains []vtypes.FailureDomain, vcenters map[string]vtypes.VCenter) (map[string]*folder, error) {
	folders := make(map[string]*folder)
	clients := make(map[string]*vim25.Client)

	for k, v := range vcenters {
		client, _, logout, err := vsphere.CreateVSphereClients(context.TODO(), v.Server, v.Username, v.Password)
		defer logout()
		if err != nil {
			return nil, err
		}
		clients[k] = client
	}

	for i, fd := range failureDomains {
		tempFolder := new(folder)
		tempFolder.Datacenter = fd.Topology.Datacenter
		tempFolder.Name = fd.Topology.Folder
		key := fmt.Sprintf("%s-%s", tempFolder.Datacenter, tempFolder.Name)

		// Only if the folder is empty do we create a folder resource
		// If a folder has been provided it means that it already exists
		// and it is to be used.
		if tempFolder.Name == "" {
			tempFolder.Name = infraID
			failureDomains[i].Topology.Folder = infraID
			folders[key] = tempFolder
		} else {
			if !folderExist(tempFolder.Name, clients[fd.Server]) {
				folders[key] = tempFolder
			}
		}
	}
	return folders, nil
}

func folderExist(folderPath string, client *vim25.Client) bool {
	ctx, timeout := context.WithTimeout(context.TODO(), 60*time.Second)
	defer timeout()
	finder := find.NewFinder(client)

	folder, err := finder.Folder(ctx, folderPath)
	if err != nil {
		return false
	}

	if folder.InventoryPath == folderPath {
		return true
	}

	return false
}
func convertVCentersToMap(values []vtypes.VCenter) map[string]vtypes.VCenter {
	vcenterMap := make(map[string]vtypes.VCenter)
	for _, v := range values {
		vcenterMap[v.Server] = v
	}
	return vcenterMap
}
