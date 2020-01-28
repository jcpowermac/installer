
package vsphere

import (
	"context"
	"net/url"
	"time"

	"github.com/openshift/installer/pkg/types/vsphere"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/vapi/rest"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/soap"
)
// Session struct contains the vCenter SOAP and REST clients
type Session struct {
	Vim25Client *vim25.Client
	RestClient  *rest.Client
}

// GetSession - creates the Session struct
func GetSession(ctx context.Context, p *vsphere.Platform) (*Session, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	u, err := soap.ParseURL(p.VCenter)
	if err != nil {
		return nil, err
	}
	u.User = url.UserPassword(p.Username, p.Username)

	// false in this method disables insecure
	// We do not allow insecure connections
	c, err := govmomi.NewClient(ctx, u, false)

	if err != nil {
		return nil, err
	}

	restClient := rest.NewClient(c.Client)
	err = restClient.Login(ctx, u.User)
	if err != nil {
		return nil, err 
	}

	return &Session{
		Vim25Client: c.Client,
		RestClient:  restClient,
	}, nil
}
