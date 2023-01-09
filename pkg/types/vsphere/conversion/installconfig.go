package conversion

import (
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/vsphere"
	"github.com/sirupsen/logrus"
)

func ConvertInstallConfig(config *types.InstallConfig) error {
	platform := config.Platform.VSphere

	if len(platform.FailureDomains) > 0 {
		if len(platform.VCenters) == 0 {
			// Should we be doing this or VCenters[] should be filled in???
			logrus.Warn("The VCenters[] field should be populated instead")

			// Should we be allowing this????
			platform.VCenters = make([]vsphere.VCenter, 1)
			platform.VCenters[0].Server = platform.DeprecatedVCenter
			platform.VCenters[0].Username = platform.DeprecatedUsername
			platform.VCenters[0].Password = platform.DeprecatedPassword
			platform.VCenters[0].Port = 443

			vcenter := &platform.VCenters[0]

			vcenter.Datacenters = append(platform.VCenters[0].Datacenters, platform.DeprecatedDatacenter)
		} else if len(platform.VCenters) > 0 {
			if !isDeprecatedFieldsEmpty(platform) {
				logrus.Warn("something something, deprecated platform fields non-empty, will not be used with FailureDomain and VCenters")
			}
		}
	} else {
		// non-zonal installation
		platform.VCenters = make([]vsphere.VCenter, 1)
		platform.VCenters[0].Server = platform.DeprecatedVCenter
		platform.VCenters[0].Username = platform.DeprecatedUsername
		platform.VCenters[0].Password = platform.DeprecatedPassword
		platform.VCenters[0].Port = 443

		vcenter := &platform.VCenters[0]

		vcenter.Datacenters = append(platform.VCenters[0].Datacenters, platform.DeprecatedDatacenter)

		platform.FailureDomains = make([]vsphere.FailureDomain, 1)
		platform.FailureDomains[0].Name = "generated-failure-domain"
		platform.FailureDomains[0].Server = platform.VCenters[0].Server
		platform.FailureDomains[0].Region = "generated-region"
		platform.FailureDomains[0].Zone = "generated-zone"

		platform.FailureDomains[0].Topology.Datacenter = platform.DeprecatedDatacenter
		platform.FailureDomains[0].Topology.ResourcePool = platform.DeprecatedResourcePool
		platform.FailureDomains[0].Topology.ComputeCluster = platform.DeprecatedCluster

		platform.FailureDomains[0].Topology.Networks = make([]string, 1)
		platform.FailureDomains[0].Topology.Networks[0] = platform.DeprecatedNetwork
		platform.FailureDomains[0].Topology.Datastore = platform.DeprecatedDefaultDatastore
		platform.FailureDomains[0].Topology.Folder = platform.DeprecatedFolder
	}

	return nil
}

func isDeprecatedFieldsEmpty(platform *vsphere.Platform) bool {
	if platform.DeprecatedVCenter == "" || platform.DeprecatedUsername == "" || platform.DeprecatedPassword == "" ||
		platform.DeprecatedDatacenter == "" || platform.DeprecatedCluster == "" || platform.DeprecatedFolder == "" ||
		platform.DeprecatedResourcePool == "" || platform.DeprecatedNetwork == "" {
		return true
	}
	return false
}
