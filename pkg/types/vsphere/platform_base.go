package vsphere

import (
	configv1 "github.com/openshift/api/config/v1"
)

// Platform stores any global configuration used for vsphere platforms.
type PlatformBase struct {
	// VCenter is the domain name or IP address of the vCenter.
	// Deprecated: Use VCenters.Server
	DeprecatedVCenter string `json:"vCenter,omitempty"`
	// Username is the name of the user to use to connect to the vCenter.
	// Deprecated: Use VCenters.Username
	DeprecatedUsername string `json:"username,omitempty"`
	// Password is the password for the user to use to connect to the vCenter.
	// Deprecated: Use VCenters.Password
	DeprecatedPassword string `json:"password,omitempty"`
	// Datacenter is the name of the datacenter to use in the vCenter.
	// Deprecated: Use FailureDomains.Topology.Datacenter
	DeprecatedDatacenter string `json:"datacenter,omitempty"`
	// DefaultDatastore is the default datastore to use for provisioning volumes.
	// Deprecated: Use FailureDomains.Topology.Datastore
	DeprecatedDefaultDatastore string `json:"defaultDatastore,omitempty"`
	// Folder is the absolute path of the folder that will be used and/or created for
	// virtual machines. The absolute path is of the form /<datacenter>/vm/<folder>/<subfolder>.
	// +kubebuilder:validation:Pattern=`^/.*?/vm/.*?`
	// +optional
	// Deprecated: Use FailureDomains.Topology.Folder
	DeprecatedFolder string `json:"folder,omitempty"`
	// Cluster is the name of the cluster virtual machines will be cloned into.
	// Deprecated: Use FailureDomains.Topology.Cluster
	DeprecatedCluster string `json:"cluster,omitempty"`
	// ResourcePool is the absolute path of the resource pool where virtual machines will be
	// created. The absolute path is of the form /<datacenter>/host/<cluster>/Resources/<resourcepool>.
	// Deprecated: Use FailureDomains.Topology.ResourcePool
	DeprecatedResourcePool string `json:"resourcePool,omitempty"`
	// ClusterOSImage overrides the url provided in rhcos.json to download the RHCOS OVA
	ClusterOSImage string `json:"clusterOSImage,omitempty"`

	// DeprecatedAPIVIP is the virtual IP address for the api endpoint
	// Deprecated: Use APIVIPs
	//
	// +kubebuilder:validation:format=ip
	// +optional
	DeprecatedAPIVIP string `json:"apiVIP,omitempty"`

	// APIVIPs contains the VIP(s) for the api endpoint. In dual stack clusters
	// it contains an IPv4 and IPv6 address, otherwise only one VIP
	//
	// +kubebuilder:validation:MaxItems=2
	// +kubebuilder:validation:UniqueItems=true
	// +kubebuilder:validation:Format=ip
	// +optional
	APIVIPs []string `json:"apiVIPs,omitempty"`

	// DeprecatedIngressVIP is the virtual IP address for ingress
	// Deprecated: Use IngressVIPs
	//
	// +kubebuilder:validation:format=ip
	// +optional
	DeprecatedIngressVIP string `json:"ingressVIP,omitempty"`

	// IngressVIPs contains the VIP(s) for ingress. In dual stack clusters it
	// contains an IPv4 and IPv6 address, otherwise only one VIP
	//
	// +kubebuilder:validation:MaxItems=2
	// +kubebuilder:validation:UniqueItems=true
	// +kubebuilder:validation:Format=ip
	// +optional
	IngressVIPs []string `json:"ingressVIPs,omitempty"`

	// DefaultMachinePlatform is the default configuration used when
	// installing on VSphere for machine pools which do not define their own
	// platform configuration.
	// +optional
	DefaultMachinePlatform *MachinePool `json:"defaultMachinePlatform,omitempty"`
	// Network specifies the name of the network to be used by the cluster.
	// Deprecated: Use FailureDomains.Topology.Network
	DeprecatedNetwork string `json:"network,omitempty"`
	// DiskType is the name of the disk provisioning type,
	// valid values are thin, thick, and eagerZeroedThick. When not
	// specified, it will be set according to the default storage policy
	// of vsphere.
	DiskType DiskType `json:"diskType,omitempty"`
	// VCenters holds the connection details for services to communicate with vCenter.
	// Currently only a single vCenter is supported.
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	VCenters []VCenter `json:"vcenters,omitempty"`
	// FailureDomains holds the VSpherePlatformFailureDomainSpec which contains
	// the definition of region, zone and the vCenter topology.
	// If this is omitted failure domains (regions and zones) will not be used.
	// +kubebuilder:validation:Optional
	FailureDomains []FailureDomain `json:"failureDomains,omitempty"`

	// LoadBalancer defines how the load balancer used by the cluster is configured.
	// LoadBalancer is available in TechPreview.
	// +optional
	LoadBalancer *configv1.VSpherePlatformLoadBalancer `json:"loadBalancer,omitempty"`
}
