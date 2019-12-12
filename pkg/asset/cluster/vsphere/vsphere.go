// Package vsphere extracts vsphere metadata from install configurations.
package vsphere

import (
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/vsphere"
)

// Metadata converts an install configuration to ovirt metadata.
func Metadata(config *types.InstallConfig) *vsphere.Metadata {
	return &vsphere.Metadata{
// TODO: add here
	}
}
