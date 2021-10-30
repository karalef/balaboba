package balaboba

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"
)

var api = url.URL{
	Scheme: "https",
	Host:   "zeapi.yandex.net",
	Path:   "/lab/api/yalm/",
}

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

func (c *Client) do(path string, data, dst interface{}) error {
	u := api
	u.Path += path

	req := http.Request{
		Method:     http.MethodGet,
		URL:        &u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Host:       u.Host,
	}
	if data != nil {
		buf := bytes.NewBuffer(nil)
		err := json.NewEncoder(buf).Encode(data)
		if err != nil {
			return err
		}
		req.Body = io.NopCloser(buf)
		req.Header.Set("Content-Type", "application/json")
		if dst != nil {
			req.Method = http.MethodPost
		}
	}

	resp, err := c.httpClient.Do(&req)
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
