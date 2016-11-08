package openid

import (
	"testing"
)

func TestDiscoverWithYadis(t *testing.T) {
	// They all redirect to the same XRDS document
	expectOpIDErr(t, "example.com/xrds",
		"foo", "bar", "", false)
	expectOpIDErr(t, "http://example.com/xrds",
		"foo", "bar", "", false)
	expectOpIDErr(t, "http://example.com/xrds-loc",
		"foo", "bar", "", false)
	expectOpIDErr(t, "http://example.com/xrds-meta",
		"foo", "bar", "", false)
}

func TestDiscoverWithHtml(t *testing.T) {
	// Yadis discovery will fail, and fall back to html.
	expectOpIDErr(t, "http://example.com/html",
		"example.com/openid", "bar-name", "http://example.com/html",
		false)
	// The first url redirects to a different URL. The redirected-to
	// url should be used as claimedID.
	expectOpIDErr(t, "http://example.com/html-redirect",
		"example.com/openid", "bar-name", "http://example.com/html",
		false)
}

func TestDiscoverBadUrl(t *testing.T) {
	expectOpIDErr(t, "http://example.com/404", "", "", "", true)
}

func expectOpIDErr(t *testing.T, uri, exOpEndpoint, exOpLocalID, exClaimedID string, exErr bool) {
	item, err := testInstance.Discover(uri)
	if (err != nil) != exErr {
		t.Errorf("Unexpected error: '%s'", err)
	} else {
		if item.OpEndpoint != exOpEndpoint {
			t.Errorf("Extracted Endpoint does not match: Exepect %s, Got %s",
				exOpEndpoint, item.OpEndpoint)
		}
		if item.OpLocalID != exOpLocalID {
			t.Errorf("Extracted LocalId does not match: Exepect %s, Got %s",
				exOpLocalID, item.OpLocalID)
		}
		if item.ClaimedID != exClaimedID {
			t.Errorf("Extracted ClaimedID does not match: Exepect %s, Got %s",
				exClaimedID, item.ClaimedID)
		}
	}
}
