package client

import "net/http"

type ClientOption func(c *Client)

func WithDomain(domain string) ClientOption {
	return func(c *Client) {
		c.root = domain
	}
}

func WithCustomHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.Client = client
	}
}
