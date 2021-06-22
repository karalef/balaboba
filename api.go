package balaboba

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

const api = "https://zeapi.yandex.net/lab/api/yalm/"
const dialTimeout = time.Duration(time.Second * 30)

var header = http.Header{
	"Content-Type": {"application/json"},
}

// New makes new balaboba api client.
func New() *Client {
	// using dialer because yandex blocks simple requests.
	d := net.Dialer{
		Timeout: dialTimeout,
	}
	return &Client{
		c: http.Client{
			Timeout: dialTimeout,
			Transport: &http.Transport{
				DialContext:         d.DialContext,
				TLSHandshakeTimeout: dialTimeout,
			},
		},
		d: &d,
	}
}

// Client is Yandex Balaboba client.
type Client struct {
	d *net.Dialer
	c http.Client
}

func (c *Client) do(req *http.Request, dst interface{}) error {
	resp, err := c.c.Do(req)
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

func (c *Client) get(path string, dst interface{}) error {
	req, err := http.NewRequest(http.MethodGet, api+path, nil)
	if err != nil {
		return err
	}
	return c.do(req, dst)
}

func (c *Client) post(path string, dst interface{}, data interface{}) error {
	buf := bytes.NewBuffer(nil)
	err := json.NewEncoder(buf).Encode(data)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, api+path, buf)
	if err != nil {
		return err
	}
	req.Header = header

	return c.do(req, dst)
}

func (c *Client) options(path string) error {
	req, err := http.NewRequest(http.MethodPost, api+path, nil)
	if err != nil {
		return err
	}

	return c.do(req, nil)
}
