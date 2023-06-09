//go:build nested

package vsphere

// Platform stores any global configuration used for vsphere platforms.
type Platform struct {
	PlatformBase
	Nested Nested `json:"nested,omitempty"`
}

type Nested struct {
	Version string `json:"version"`

	ContentLibrary string `json:"contentLibrary"`
}
