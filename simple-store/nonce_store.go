package simplestore

import (
	"errors"
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/rylio/openid-go"
)

var maxNonceAge = flag.Duration("openid-max-nonce-age",
	60*time.Second,
	"Maximum accepted age for openid nonces. The bigger, the more"+
		"memory is needed to store used nonces.")

type SimpleNonceStore struct {
	store map[string][]openid.NonceItem
	mutex *sync.Mutex
}

func NewSimpleNonceStore() *SimpleNonceStore {
	return &SimpleNonceStore{store: make(map[string][]openid.NonceItem), mutex: &sync.Mutex{}}
}

func (d *SimpleNonceStore) Accept(endpoint string, nonce openid.NonceItem) error {

	now := time.Now()
	diff := now.Sub(nonce.Time)
	if diff > *maxNonceAge {
		return fmt.Errorf("Nonce too old: %ds", diff.Seconds())
	}

	// Meh.. now we have to use a mutex, to protect that map from
	// concurrent access. Could put a go routine in charge of it
	// though.
	d.mutex.Lock()
	defer d.mutex.Unlock()

	if nonces, hasOp := d.store[endpoint]; hasOp {
		// Delete old nonces while we are at it.
		newNonces := []openid.NonceItem{nonce}
		for _, n := range nonces {
			if n == nonce {
				// If return early, just ignore the filtered list
				// we have been building so far...
				return errors.New("Nonce already used")
			}
			if now.Sub(n.Time) < *maxNonceAge {
				newNonces = append(newNonces, n)
			}
		}
		d.store[endpoint] = newNonces
	} else {
		d.store[endpoint] = []openid.NonceItem{nonce}
	}
	return nil
}
