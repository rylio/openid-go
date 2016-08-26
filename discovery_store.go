package openid

type DiscoveryItem struct {
	OpEndpoint string
	OpLocalID  string
	ClaimedID  string
}

type DiscoveryStore interface {
	Put(id string, info *DiscoveryItem) error
	// Return a discovered info, or nil.
	Get(id string) (*DiscoveryItem, error)
}
