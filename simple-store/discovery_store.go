package simplestore

import (
	"sync"

	"github.com/rylio/openid-go"
)

type DiscoveryStore struct {
	cache map[string]*openid.DiscoveryItem
	mutex *sync.Mutex
}

func NewDiscoveryStore() *DiscoveryStore {
	return &DiscoveryStore{cache: make(map[string]*openid.DiscoveryItem), mutex: &sync.Mutex{}}
}

func (s *DiscoveryStore) Put(id string, item *openid.DiscoveryItem) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache[id] = item
	return nil
}

func (s *DiscoveryStore) Get(id string) (*openid.DiscoveryItem, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if info, has := s.cache[id]; has {
		return info, nil
	}
	return nil, nil
}
