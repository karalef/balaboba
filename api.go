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

// Timeout
const Timeout = 15

var api = url.URL{
	Scheme: "https",
	Host:   "zeapi.yandex.net",
	Path:   "/lab/api/yalm/",
}

// New makes new balaboba api client.
func New() *Client {
	// using dialer because yandex blocks simple requests.
	d := net.Dialer{
		Timeout: time.Second * Timeout,
	}
	return &Client{
		c: http.Client{
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
	c http.Client
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
	} else if dst == nil {
		req.Method = http.MethodOptions
	}

	resp, err := c.c.Do(&req)
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
