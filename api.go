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

const apiurl = "https://yandex.ru/lab/api/yalm/"

// ClientRus var.
var ClientRus = New(Rus)

// ClientEng var.
var ClientEng = New(Eng)

// New makes new balaboba api client.
func New(lang Lang) *Client {
	d := net.Dialer{
		Timeout: 20 * time.Second,
	}
	return &Client{
		lang: lang,
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
	lang       Lang
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
		method = http.MethodPost
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

// Lang represents balaboba language.
type Lang uint8

// available languages.
const (
	Rus Lang = iota
	Eng
)
