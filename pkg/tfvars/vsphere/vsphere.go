package vsphere

import (
	"encoding/json"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"strings"

	"github.com/pkg/errors"

	machineapi "github.com/openshift/api/machine/v1beta1"
	"github.com/openshift/installer/pkg/tfvars/internal/cache"
	"github.com/openshift/installer/pkg/types/vsphere"
	vspheretypes "github.com/openshift/installer/pkg/types/vsphere"
)

type config struct {
	VSphereURL        string           `json:"vsphere_url"`
	VSphereUsername   string           `json:"vsphere_username"`
	VSpherePassword   string           `json:"vsphere_password"`
	MemoryMiB         int64            `json:"vsphere_control_plane_memory_mib"`
	DiskGiB           int32            `json:"vsphere_control_plane_disk_gib"`
	NumCPUs           int32            `json:"vsphere_control_plane_num_cpus"`
	NumCoresPerSocket int32            `json:"vsphere_control_plane_cores_per_socket"`
	Cluster           string           `json:"vsphere_cluster"`
	ResourcePool      string           `json:"vsphere_resource_pool"`
	Datacenter        string           `json:"vsphere_datacenter"`
	Datastore         string           `json:"vsphere_datastore"`
	Folder            string           `json:"vsphere_folder"`
	Network           string           `json:"vsphere_network"`
	Template          string           `json:"vsphere_template"`
	OvaFilePath       string           `json:"vsphere_ova_filepath"`
	PreexistingFolder bool             `json:"vsphere_preexisting_folder"`
	DiskType          vsphere.DiskType `json:"vsphere_disk_type"`

	VCenters *[]vspheretypes.VCenter `json:"vsphere_vcenters"`
}

// TFVarsSources contains the parameters to be converted into Terraform variables
type TFVarsSources struct {
	ControlPlaneConfigs []*machineapi.VSphereMachineProviderSpec
	Username            string
	Password            string
	Cluster             string
	ImageURL            string
	PreexistingFolder   bool
	DiskType            vsphere.DiskType
	InstallConfig       *installconfig.InstallConfig
}

//TFVars generate vSphere-specific Terraform variables
func TFVars(sources TFVarsSources) ([]byte, error) {

	platform := sources.InstallConfig.Config.VSphere

	// TODO: This needs to be the slice
	controlPlaneConfig := sources.ControlPlaneConfigs

	cachedImage, err := cache.DownloadImageFile(sources.ImageURL)
	if err != nil {
		return nil, errors.Wrap(err, "failed to use cached vsphere image")
	}

	// The vSphere provider needs the relativepath of the folder,
	// so get the relPath from the absolute path. Absolute path is always of the form
	// /<datacenter>/vm/<folder_path> so we can split on "vm/".
	folderRelPath := strings.SplitAfterN(controlPlaneConfig[0].Workspace.Folder, "vm/", 2)[1]

	// TODO: do not break existing install
	cfg := &config{
		VSphereURL:        controlPlaneConfig[0].Workspace.Server,
		VSphereUsername:   sources.Username,
		VSpherePassword:   sources.Password,
		MemoryMiB:         controlPlaneConfig[0].MemoryMiB,
		DiskGiB:           controlPlaneConfig[0].DiskGiB,
		NumCPUs:           controlPlaneConfig[0].NumCPUs,
		NumCoresPerSocket: controlPlaneConfig[0].NumCoresPerSocket,
		Cluster:           sources.Cluster,
		ResourcePool:      controlPlaneConfig[0].Workspace.ResourcePool,
		Datacenter:        controlPlaneConfig[0].Workspace.Datacenter,
		Datastore:         controlPlaneConfig[0].Workspace.Datastore,
		Folder:            folderRelPath,
		Network:           controlPlaneConfig[0].Network.Devices[0].NetworkName,

		// TODO: change me...
		Template:          controlPlaneConfig[0].Template,
		OvaFilePath:       cachedImage,
		PreexistingFolder: sources.PreexistingFolder,
		DiskType:          sources.DiskType,
		VCenters:          &platform.VCenters,
	}

	return json.MarshalIndent(cfg, "", "  ")
}
