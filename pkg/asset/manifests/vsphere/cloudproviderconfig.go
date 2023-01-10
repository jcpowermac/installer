package vsphere

import (
	"bytes"
	"fmt"
	"gopkg.in/ini.v1"
	"strconv"
	"strings"

	"github.com/go-yaml/yaml"
	cloudconfig "k8s.io/cloud-provider-vsphere/pkg/common/config"

	vspheretypes "github.com/openshift/installer/pkg/types/vsphere"
)

const (
	regionTagCategory = "openshift-region"
	zoneTagCategory   = "openshift-zone"
)

func appendTagCategory(tagCategory string, tagCategories []string) []string {
	tagDefined := false
	for _, regionTagCategory := range tagCategories {
		if regionTagCategory == tagCategory {
			tagDefined = true
			break
		}
	}
	if tagDefined == false {
		return append(tagCategories, tagCategory)
	}
	return tagCategories
}

// MultiZoneYamlCloudProviderConfig generates the yaml out of tree cloud provider config for the vSphere platform.
func MultiZoneYamlCloudProviderConfig(p *vspheretypes.Platform) (string, error) {
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

// IniCloudProviderConfig generates external CCM ini configuration
// for the vSphere platform. p is the vSphere platform struct.
func IniCloudProviderConfig(p *vspheretypes.Platform) (string, error) {
	buf := new(bytes.Buffer)
	cfg := ini.Empty()

	global, err := setOrGetSection(cfg, "Global")
	if err != nil {
		return "", err
	}
	_, err = setOrGetKeyValue(global, "secret-name", "vsphere-creds", true)
	if err != nil {
		return "", err
	}
	_, err = setOrGetKeyValue(global, "secret-namespace", "kube-system", true)
	if err != nil {
		return "", err
	}
	_, err = setOrGetKeyValue(global, "insecure-flag", "1", true)
	if err != nil {
		return "", err
	}

	if err := setVirtualCenters(cfg, p); err != nil {
		return "", err
	}
	if err = setDatacentersFromFailureDomains(cfg, p); err != nil {
		return "", err
	}
	// TODO: does this make sense to not enable for len(failuredomains) == 1
	// TODO: thinking in the scenario of generated fds
	if len(p.FailureDomains) > 1 {
		if err := setLabelsSection(cfg); err != nil {
			return "", err
		}
	}

	_, err = cfg.WriteTo(buf)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func setOrGetSection(cfg *ini.File, name string) (*ini.Section, error) {
	if cfg.HasSection(name) {
		return cfg.GetSection(name)
	} else {
		return cfg.NewSection(name)
	}
}
func setOrGetKeyValue(section *ini.Section, name, value string, overwrite bool) (*ini.Key, error) {
	var err error
	var key *ini.Key
	if !section.HasKey(name) {
		key, err = section.NewKey(name, value)
	} else {
		key, err = section.GetKey(name)
		if err != nil {
			return nil, err
		}
		if overwrite {
			key.SetValue(value)
		}
	}
	return key, err
}
func setVirtualCenters(cfg *ini.File, p *vspheretypes.Platform) error {
	for _, vcenter := range p.VCenters {
		vcenterSectionName := fmt.Sprintf("VirtualCenter \"%s\"", vcenter.Server)
		vCenterSection, err := setOrGetSection(cfg, vcenterSectionName)
		if err != nil {
			return fmt.Errorf("could not get or set VirtualCenter section: %w", err)
		}

		if err := setVCenterPortKey(vCenterSection, vcenter.Port); err != nil {
			return fmt.Errorf("could not set VirtualCenters port value: %w", err)
		}

		datacenters := strings.Join(vcenter.Datacenters[:], ",")
		_, err = setOrGetKeyValue(vCenterSection, "datacenters", datacenters, true)
		if err != nil {
			return fmt.Errorf("could not get or set the datacenters key value: %w", err)
		}
	}
	return nil
}
func setVCenterPortKey(vCenterSection *ini.Section, port int32) error {
	_, err := setOrGetKeyValue(vCenterSection, "port", strconv.FormatInt(int64(port), 10), true)
	return err
}
func setDatacentersFromFailureDomains(cfg *ini.File, p *vspheretypes.Platform) error {
	for _, fd := range p.FailureDomains {
		vcenterSectionName := fmt.Sprintf("VirtualCenter \"%s\"", fd.Server)
		vCenterSection, err := setOrGetSection(cfg, vcenterSectionName)
		if err != nil {
			return err
		}

		err = setVCenterDatacentersKey(vCenterSection, fd.Topology.Datacenter)
		if err != nil {
			return err
		}
	}
	return nil
}
func setVCenterDatacentersKey(vCenterSection *ini.Section, datacenter string) error {
	// Get existing key (don't overwrite) or get new key
	key, err := setOrGetKeyValue(vCenterSection, "datacenters", datacenter, false)
	if err != nil {
		return err
	}

	datacenters := key.String()

	if !strings.Contains(datacenters, datacenter) {
		datacenters = fmt.Sprintf("%s,%s", datacenters, datacenter)
		key.SetValue(datacenters)
	}

	return nil
}
func setLabelsSection(cfg *ini.File) error {
	labels := map[string]string{"region": regionTagCategory, "zone": zoneTagCategory}

	labelsSection, err := setOrGetSection(cfg, "Labels")
	if err != nil {
		return err
	}

	for k, v := range labels {
		_, err = setOrGetKeyValue(labelsSection, k, v, true)
		if err != nil {
			return err
		}
	}

	return nil
}
