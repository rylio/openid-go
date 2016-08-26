package openid

import "time"

type NonceStore interface {
	// Returns nil if accepted, an error otherwise.
	Accept(endpoint string, nonce NonceItem) error
}

type NonceItem struct {
	Time  time.Time
	Nonce string
}
