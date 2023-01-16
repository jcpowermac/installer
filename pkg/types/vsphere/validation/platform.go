package validation

import (
	"fmt"
	"regexp"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/openshift/installer/pkg/types/vsphere"
	"github.com/openshift/installer/pkg/validate"
)

// ValidatePlatform checks that the specified platform is valid.
func ValidatePlatform(p *vsphere.Platform, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	if len(p.VCenters) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("vcenters"), "must be defined"))
		return allErrs
	} else {
		allErrs = append(allErrs, validateVCenters(p, fldPath.Child("vcenters"))...)
	}
	if len(p.FailureDomains) == 0 {
		allErrs = append(allErrs, field.Required(fldPath.Child("failureDomains"), "must be defined"))
	} else {
		allErrs = append(allErrs, validateFailureDomains(p, fldPath.Child("failureDomains"))...)
	}
	// diskType is optional, but if provided should pass validation
	if len(p.DiskType) != 0 {
		allErrs = append(allErrs, validateDiskType(p, fldPath)...)
	}

	return allErrs
}

func validateVCenters(p *vsphere.Platform, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	if len(p.VCenters) > 1 {
		return field.ErrorList{field.TooMany(fldPath, len(p.VCenters), 1)}
	}

	for _, vCenter := range p.VCenters {
		if len(vCenter.Server) == 0 {
			allErrs = append(allErrs, field.Required(fldPath.Child("server"), "must be the domain name or IP address of the vCenter"))
		} else {
			if err := validate.Host(vCenter.Server); err != nil {
				allErrs = append(allErrs, field.Invalid(fldPath.Child("server"), vCenter.Server, "must be the domain name or IP address of the vCenter"))
			}
		}
		if len(vCenter.Username) == 0 {
			allErrs = append(allErrs, field.Required(fldPath.Child("username"), "must specify the username"))
		}
		if len(vCenter.Password) == 0 {
			allErrs = append(allErrs, field.Required(fldPath.Child("password"), "must specify the password"))
		}
		if len(vCenter.Datacenters) == 0 {
			allErrs = append(allErrs, field.Required(fldPath.Child("datacenters"), "must specify at least one datacenter"))
		}
	}
	return allErrs
}

func validateFailureDomains(p *vsphere.Platform, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	topologyFld := fldPath.Child("topology")
	var associatedVCenter *vsphere.VCenter
	for _, failureDomain := range p.FailureDomains {
		if len(failureDomain.Name) == 0 {
			allErrs = append(allErrs, field.Required(fldPath.Child("name"), "must specify the name"))
		}
		if len(failureDomain.Server) > 0 {
			for _, vcenter := range p.VCenters {
				if vcenter.Server == failureDomain.Server {
					associatedVCenter = &vcenter
					break
				}
			}
			if associatedVCenter == nil {
				allErrs = append(allErrs, field.Invalid(fldPath.Child("server"), failureDomain.Server, "server does not exist in vcenters"))
			}
		} else {
			allErrs = append(allErrs, field.Required(fldPath.Child("server"), "must specify a vCenter server"))
		}

		if len(failureDomain.Zone) == 0 {
			allErrs = append(allErrs, field.Required(fldPath.Child("zone"), "must specify zone tag value"))
		}

		if len(failureDomain.Region) == 0 {
			allErrs = append(allErrs, field.Required(fldPath.Child("region"), "must specify region tag value"))
		}

		if len(failureDomain.Topology.Datacenter) == 0 {
			allErrs = append(allErrs, field.Required(topologyFld.Child("datacenter"), "must specify a datacenter"))
		}

		if len(failureDomain.Topology.Datastore) == 0 {
			allErrs = append(allErrs, field.Required(topologyFld.Child("datastore"), "must specify a datastore"))
		} else {

			// todo: jcallen: should we generate a datastore path?
			datastore := failureDomain.Topology.Datastore
			datastorePathRegexp := regexp.MustCompile(`^/(.*?)/datastore/(.*?)$`)
			datastorePathParts := datastorePathRegexp.FindStringSubmatch(datastore)
			if len(datastorePathParts) < 3 {
				return append(allErrs, field.Invalid(topologyFld.Child("datastore"), datastore, "full path of datastore must be provided in format /<datacenter/datastore/<datastore>"))
			}

			if !strings.Contains(failureDomain.Topology.Datastore, failureDomain.Topology.Datacenter) {
				return append(allErrs, field.Invalid(topologyFld.Child("datastore"), failureDomain.Topology.Datastore, "the datastore defined does not exist in the correct datacenter"))
			}
		}

		if len(failureDomain.Topology.Networks) == 0 {
			allErrs = append(allErrs, field.Required(topologyFld.Child("networks"), "must specify a network"))
		}
		// Folder in failuredomain is optional
		if len(failureDomain.Topology.Folder) != 0 {
			folderPathRegexp := regexp.MustCompile(`^/(.*?)/vm/(.*?)$`)
			folderPathParts := folderPathRegexp.FindStringSubmatch(failureDomain.Topology.Folder)
			if len(folderPathParts) < 3 {
				return append(allErrs, field.Invalid(topologyFld.Child("folder"), failureDomain.Topology.Folder, "full path of folder must be provided in format /<datacenter>/vm/<folder>"))
			}

			if !strings.Contains(failureDomain.Topology.Folder, failureDomain.Topology.Datacenter) {
				return append(allErrs, field.Invalid(topologyFld.Child("folder"), failureDomain.Topology.Folder, "the folder defined does not exist in the correct datacenter"))
			}
		}
		// ResourcePool in failuredomain is optional
		if len(failureDomain.Topology.ResourcePool) != 0 {
			folderPathRegexp := regexp.MustCompile(`^/(.*?)/host/(.*?)/Resources(.*?)$`)
			folderPathParts := folderPathRegexp.FindStringSubmatch(failureDomain.Topology.Folder)
			if len(folderPathParts) < 3 {
				return append(allErrs, field.Invalid(topologyFld.Child("folder"), failureDomain.Topology.Folder, "full path of folder must be provided in format /<datacenter>/vm/<folder>"))
			}

			if !strings.Contains(failureDomain.Topology.Folder, failureDomain.Topology.Datacenter) {
				return append(allErrs, field.Invalid(topologyFld.Child("folder"), failureDomain.Topology.Folder, "the folder defined does not exist in the correct datacenter"))
			}
		}

		if len(failureDomain.Topology.ComputeCluster) == 0 {
			allErrs = append(allErrs, field.Required(topologyFld.Child("computeCluster"), "must specify a computeCluster"))
		} else {
			computeCluster := failureDomain.Topology.ComputeCluster
			clusterPathRegexp := regexp.MustCompile(`^/(.*?)/host/(.*?)$`)
			clusterPathParts := clusterPathRegexp.FindStringSubmatch(computeCluster)
			if len(clusterPathParts) < 3 {
				return append(allErrs, field.Invalid(topologyFld.Child("computeCluster"), computeCluster, "full path of compute cluster must be provided in format /<datacenter>/host/<cluster>"))
			}
			datacenterName := clusterPathParts[1]

			if len(failureDomain.Topology.Datacenter) != 0 && datacenterName != failureDomain.Topology.Datacenter {
				return append(allErrs, field.Invalid(topologyFld.Child("computeCluster"), computeCluster, fmt.Sprintf("compute cluster must be in datacenter %s", failureDomain.Topology.Datacenter)))
			}
		}
	}

	return allErrs
}

// validateDiskType checks that the specified diskType is valid
func validateDiskType(p *vsphere.Platform, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	validDiskTypes := sets.NewString(string(vsphere.DiskTypeThin), string(vsphere.DiskTypeThick), string(vsphere.DiskTypeEagerZeroedThick))
	if !validDiskTypes.Has(string(p.DiskType)) {
		errMsg := fmt.Sprintf("diskType must be one of %v", validDiskTypes.List())
		allErrs = append(allErrs, field.Invalid(fldPath.Child("diskType"), p.DiskType, errMsg))
	}

	return allErrs
}
