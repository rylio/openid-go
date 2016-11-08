package simplestore

import (
	"testing"

	openid "github.com/rylio/openid-go"
)

func TestDiscoveryStore(t *testing.T) {
	dc := NewDiscoveryStore()

	// Put some initial values
	dc.Put("foo", &openid.DiscoveryItem{OpEndpoint: "a", OpLocalID: "b", ClaimedID: "c"})

	// Make sure we can retrieve them
	if di, _ := dc.Get("foo"); di == nil {
		t.Errorf("Expected a result, got nil")
	} else if di.OpEndpoint != "a" || di.OpLocalID != "b" || di.ClaimedID != "c" {
		t.Errorf("Expected a b c, got %v %v %v", di.OpEndpoint, di.OpLocalID, di.ClaimedID)
	}

	// Attempt to get a non-existent value
	if di, _ := dc.Get("bar"); di != nil {
		t.Errorf("Expected nil, got %v", di)
	}
}
