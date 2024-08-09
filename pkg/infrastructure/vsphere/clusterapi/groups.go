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
 3. TODO: jcallen: **** CLUSTER ID **** is wrong, it needs to be probably failure domain name
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

func addVmHostRule(ctx context.Context, session *session.Session, cluster, hostGroupName, clusterID string) error {
	clusterObj, err := session.Finder.ClusterComputeResource(ctx, cluster)
	if err != nil {
		return err
	}
	clusterConfigSpec := &types.ClusterConfigSpecEx{
		RulesSpec: []types.ClusterRuleSpec{
			{
				ArrayUpdateSpec: types.ArrayUpdateSpec{
					Operation: types.ArrayUpdateOperation("add"),
				},
				Info: &types.ClusterVmHostRuleInfo{
					ClusterRuleInfo: types.ClusterRuleInfo{
						Name:        clusterID,
						Mandatory:   types.NewBool(true),
						Enabled:     types.NewBool(true),
						UserCreated: types.NewBool(true),
					},
					VmGroupName:         clusterID,
					AffineHostGroupName: hostGroupName,
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
