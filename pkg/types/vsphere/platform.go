package vsphere

// Platform stores any global configuration used for vsphere platforms.
type Platform struct {
	// VCenter is the domain name or IP address of the vCenter.
	VCenter string `json:"vCenter"`
	// Username is the name of the user to use to connect to the vCenter.
	Username string `json:"username"`
	// Password is the password for the user to use to connect to the vCenter.
	Password string `json:"password"`
	// Datacenter is the name of the datacenter to use in the vCenter.
	Datacenter string `json:"datacenter"`
	// DefaultDatastore is the default datastore to use for provisioning volumes.
	DefaultDatastore string `json:"defaultDatastore"`
	// Network is the default network to use for vm networking
	Network string `json:"network"`
	// Template is the virtual machine or template that will be used when cloning to
	// create a new virtual machine.
	Template string `json:"template"`
	// ResourcePool is the resource pool that will be used.
	// If empty a RP will not be used.  RP will not be created
	ResourcePool string `json:"resourcePool,omitempty"`
	// Folder is the the folder that will be used or created to contain the cluster virtual machines
	Folder string `json:"folder,omitempty"`
	// APIVIP is the VIP to use for internal API communication
	APIVIP string `json:"apiVIP,omitempty"`
	// IngressVIP is the VIP to use for ingress traffic
	IngressVIP string `json:"ingressVIP,omitempty"`
	// DNSVIP is the VIP to use for internal DNS communication
	DNSVIP string `json:"dnsVIP,omitempty"`
}
