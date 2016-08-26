package simplestore

import (
	"sync"

	"github.com/rylio/openid-go"
)

type SimpleDiscoveryStore struct {
	cache map[string]*openid.DiscoveryItem
	mutex *sync.Mutex
}

func NewSimpleDiscoveryStore() *SimpleDiscoveryStore {
	return &SimpleDiscoveryStore{cache: make(map[string]*openid.DiscoveryItem), mutex: &sync.Mutex{}}
}

func (s *SimpleDiscoveryStore) Put(id string, item *openid.DiscoveryItem) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache[id] = item
	return nil
}

func (s *SimpleDiscoveryStore) Get(id string) (*openid.DiscoveryItem, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if info, has := s.cache[id]; has {
		return info, nil
	}
	return nil, nil
}
