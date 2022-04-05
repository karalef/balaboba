package balaboba

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

const apiurl = "http://zeapi.yandex.net/lab/api/yalm/"

// New makes new balaboba api client.
func New() *Client {
	d := net.Dialer{
		Timeout: 15 * time.Second,
	}
	return &Client{
		httpClient: http.Client{
			Timeout: d.Timeout,
			Transport: &http.Transport{
				DialContext:         d.DialContext,
				TLSHandshakeTimeout: d.Timeout,
			},
		},
	}
}

// Client is Yandex Balaboba client.
type Client struct {
	httpClient http.Client
}

func (c *Client) do(ctx context.Context, path string, data, dst interface{}) error {
	method := http.MethodGet
	var body io.Reader

	if data != nil {
		buf, err := json.Marshal(data)
		if err != nil {
			return err
		}
		body = bytes.NewReader(buf)
		if dst != nil {
			method = http.MethodPost
		}
	} else if dst == nil {
		method = http.MethodOptions
	}

	if ctx == nil {
		ctx = context.Background()
	}

	req, _ := http.NewRequestWithContext(ctx, method, apiurl+path, body)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s (%d)", resp.Status, resp.StatusCode)
	}
	if dst != nil {
		err = json.NewDecoder(resp.Body).Decode(dst)
	}
	resp.Body.Close()
	return err
}

// IsAvailable checks the service for availability.
func (c *Client) IsAvailable() bool {
	return c.do(nil, text3api, nil, nil) == nil
}
