package openid

import (
	"net/http"
)

type Client struct {
	urlGetter httpGetter
}

func New(client *http.Client) *Client {
	return &Client{urlGetter: &defaultGetter{client: client}}
}
