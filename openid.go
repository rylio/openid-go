package openid

import (
	"net/http"
)

type OpenID struct {
	urlGetter httpGetter
}

func NewOpenID(client *http.Client) *OpenID {
	return &OpenID{urlGetter: &defaultGetter{client: client}}
}
