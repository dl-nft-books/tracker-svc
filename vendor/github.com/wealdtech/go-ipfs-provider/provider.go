package provider

import "io"

// Provider is the interface an IPFS provider service must support.
type Provider interface {
	// Ping pings the provider to ensure it is alive.
	Ping() (bool, error)

	// List lists the files pinned on the provider.
	List() ([]*ItemStatistics, error)

	// ItemStats provides statistics for the pinned item.
	ItemStats(string) (*ItemStatistics, error)

	// ServiceStats obtains statistics for this provider.
	ServiceStats() (*SiteStatistics, error)

	// PinContent pins a provided file to this provider.
	PinContent(string, io.Reader, *ContentOpts) (string, error)

	// Pin pins an existing IPFS hash to this provider.
	Pin(string) error

	// Unpin unpins an existing IPFS hash from this provider.
	Unpin(string) error

	// GatewayURL returns a gateway URL for an IPFS hash.
	GatewayURL(string) (string, error)
}

// ContentOpts are options for storing content
type ContentOpts struct {
	// StoreInDirectory will store the content in a directory if true
	StoreInDirectory bool
}

// ItemStatistics is a structure for a piece of IPFS content.
type ItemStatistics struct {
	// Name is the name of the item.
	Name string
	// Hash is the IPFS hash of the item.
	Hash string
	// Size is the size of the item in bytes.
	Size uint64
}

// SiteStatistics is a structure for a provider service.
type SiteStatistics struct {
	// Items is the total number of items pinned.
	Items uint64
	// Size is the total size of items pinned.
	Size uint64
}
