//go:build altinfra || aro
// +build altinfra aro

package platform

import (
	"fmt"

	"github.com/openshift/installer/pkg/infrastructure"
	"github.com/openshift/installer/pkg/infrastructure/vsphere"
	awstypes "github.com/openshift/installer/pkg/types/aws"
	azuretypes "github.com/openshift/installer/pkg/types/azure"
	vspheretypes "github.com/openshift/installer/pkg/types/vsphere"
)

// ProviderForPlatform returns the stages to run to provision the infrastructure for the specified platform.
func ProviderForPlatform(platform string) (infrastructure.Provider, error) {
	switch platform {
	case awstypes.Name:
		panic("not implemented")
		return nil, nil
	case azuretypes.Name:
		panic("not implemented")
		return nil, nil
	case vspheretypes.Name:
		return vsphere.InitializeProvider(), nil
	}
	return nil, fmt.Errorf("platform %q is not supported in the altinfra Installer build", platform)
}
