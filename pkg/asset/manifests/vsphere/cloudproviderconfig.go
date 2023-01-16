package vsphere

import (
	"github.com/go-yaml/yaml"
	cloudconfig "k8s.io/cloud-provider-vsphere/pkg/common/config"

	vspheretypes "github.com/openshift/installer/pkg/types/vsphere"
)

const (
	regionTagCategory = "openshift-region"
	zoneTagCategory   = "openshift-zone"
)

// CloudProviderConfig generates the yaml out of tree cloud provider config for the vSphere platform.
func CloudProviderConfig(p *vspheretypes.Platform) (string, error) {
	vCenters := make(map[string]*cloudconfig.VirtualCenterConfigYAML)

	for _, vCenter := range p.VCenters {
		vCenterPort := int32(443)
		if vCenter.Port != 0 {
			vCenterPort = vCenter.Port
		}
		vCenterConfig := cloudconfig.VirtualCenterConfigYAML{
			VCenterIP:   vCenter.Server,
			VCenterPort: uint(vCenterPort),
			Datacenters: vCenter.Datacenters,
		}
		vCenters[vCenter.Server] = &vCenterConfig
	}

	cloudProviderConfig := cloudconfig.CommonConfigYAML{
		Global: cloudconfig.GlobalYAML{
			SecretName:      "vsphere-creds",
			SecretNamespace: "kube-system",
		},
		Vcenter: vCenters,
		Labels: cloudconfig.LabelsYAML{
			Zone:   zoneTagCategory,
			Region: regionTagCategory,
		},
	}

	cloudProviderConfigYaml, err := yaml.Marshal(cloudProviderConfig)
	if err != nil {
		return "", err
	}
	return string(cloudProviderConfigYaml), nil
}
