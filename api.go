package balaboba

import (
	"net/http"
)

const api = "https://zeapi.yandex.net/lab/api/yalm/"

// New makes new balaboba api client.
func New(c *http.Client) *Client {
	if c == nil {
		c = http.DefaultClient
	}
	return &Client{c}
}

// Client is Yandex Balaboba client.
type Client struct {
	c *http.Client
}
