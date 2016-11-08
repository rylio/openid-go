package simplestore

import (
	"testing"
	"time"

	openid "github.com/rylio/openid-go"
)

func TestNonceStore(t *testing.T) {
	now := time.Now().UTC()
	// 30 seconds ago
	now30s := now.Add(-30 * time.Second)
	// 2 minutes ago
	now2m := now.Add(-2 * time.Minute)

	//now30sStr := now30s.Format(time.RFC3339)
	//now2mStr := now2m.Format(time.RFC3339)

	ns := NewNonceStore()
	ns.MaxNonceAge = time.Minute
	var s string
	for i := 0; i < 300; i++ {
		s += "a"
	}
	reject(t, ns, "1", openid.NonceItem{now, s}) // invalid nonce

	accept(t, ns, "1", openid.NonceItem{now30s, "asd"})
	reject(t, ns, "1", openid.NonceItem{now30s, "asd"}) // same nonce
	accept(t, ns, "1", openid.NonceItem{now30s, "xxx"}) // different nonce
	reject(t, ns, "1", openid.NonceItem{now30s, "xxx"}) // different nonce again to verify storage of multiple nonces per endpoint
	accept(t, ns, "2", openid.NonceItem{now30s, "asd"}) // different endpoint

	reject(t, ns, "1", openid.NonceItem{now2m, "old"}) // too old
	reject(t, ns, "3", openid.NonceItem{now2m, "old"}) // too old
}

func accept(t *testing.T, ns *NonceStore, op string, nonce openid.NonceItem) {
	e := ns.Accept(op, nonce)
	if e != nil {
		t.Errorf("Should accept %s nonce %s", op, nonce)
	}
}

func reject(t *testing.T, ns *NonceStore, op string, nonce openid.NonceItem) {
	e := ns.Accept(op, nonce)
	if e == nil {
		t.Errorf("Should reject %s nonce %s", op, nonce)
	}
}
