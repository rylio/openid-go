package openid

import (
	"net/url"
	"strings"
)

func (oid *Client) RedirectURL(id, callbackURL, realm string) (string, error) {
	item, err := oid.Discover(id)
	if err != nil {
		return "", err
	}
	return BuildRedirectURL(item, callbackURL, realm)
}

func BuildRedirectURL(item DiscoveryItem, returnTo, realm string) (string, error) {
	values := make(url.Values)
	values.Add("openid.ns", "http://specs.openid.net/auth/2.0")
	values.Add("openid.mode", "checkid_setup")
	values.Add("openid.return_to", returnTo)

	// 9.1.  Request Parameters
	// "openid.claimed_id" and "openid.identity" SHALL be either both present or both absent.
	if len(item.ClaimedID) > 0 {
		values.Add("openid.claimed_id", item.ClaimedID)
		if len(item.OpLocalID) > 0 {
			values.Add("openid.identity", item.OpLocalID)
		} else {
			// If a different OP-Local Identifier is not specified,
			// the claimed identifier MUST be used as the value for openid.identity.
			values.Add("openid.identity", item.ClaimedID)
		}
	} else {
		// 7.3.1.  Discovered Information
		// If the end user entered an OP Identifier, there is no Claimed Identifier.
		// For the purposes of making OpenID Authentication requests, the value
		// "http://specs.openid.net/auth/2.0/identifier_select" MUST be used as both the
		// Claimed Identifier and the OP-Local Identifier when an OP Identifier is entered.
		values.Add("openid.claimed_id", "http://specs.openid.net/auth/2.0/identifier_select")
		values.Add("openid.identity", "http://specs.openid.net/auth/2.0/identifier_select")
	}

	if len(realm) > 0 {
		values.Add("openid.realm", realm)
	}

	if strings.Contains(item.OpEndpoint, "?") {
		return item.OpEndpoint + "&" + values.Encode(), nil
	}
	return item.OpEndpoint + "?" + values.Encode(), nil
}
