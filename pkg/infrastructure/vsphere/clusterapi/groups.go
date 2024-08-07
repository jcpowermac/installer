package clusterapi

import (
	"context"

	"github.com/vmware/govmomi/vim25/types"
	"sigs.k8s.io/cluster-api-provider-vsphere/pkg/session"
)

/*
TODO: jcallen:
 1. does the vm group already exist?
 2. how do we add a vm to a group?
*/
func createVmGroup(ctx context.Context, session *session.Session, cluster, clusterID string) error {
	clusterObj, err := session.Finder.ClusterComputeResource(ctx, cluster)
	if err != nil {
		return err
	}

	clusterConfigSpec := &types.ClusterConfigSpecEx{
		GroupSpec: []types.ClusterGroupSpec{
			{
				ArrayUpdateSpec: types.ArrayUpdateSpec{
					Operation: types.ArrayUpdateOperation("add"),
				},
				Info: &types.ClusterVmGroup{
					ClusterGroupInfo: types.ClusterGroupInfo{
						Name: clusterID,
					},
				},
			},
		},
	}

	task, err := clusterObj.Reconfigure(ctx, clusterConfigSpec, true)
	if err != nil {
		return err
	}

	if err := task.Wait(ctx); err != nil {
		return err
	}
	return nil
}
