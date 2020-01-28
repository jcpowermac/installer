package vsphere

import (
	"context"

	vspheretypes "github.com/openshift/installer/pkg/types/vsphere"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
)

func ValidateCreds(ssn *Session, p *vspheretypes.Platform) error {

	// What is my user account?
	// What group(s) is my user account in?
	// For the datacenter/host/cluster
	// user - role association
	// group - role association
	// what permissions do the roles contain?
	ctx := context.TODO()
	authManager := object.NewAuthorizationManager(ssn.Vim25Client)

	finder := find.NewFinder(ssn.Vim25Client)

	datacenter, err := finder.Datacenter(ctx, p.Datacenter)

	if err != nil {
		return err
	}

	permissions, err := authManager.RetrieveEntityPermissions(ctx, datacenter.Reference(), true)

	for _, p := range permissions {
		p.Principal

	}

	return nil
}
