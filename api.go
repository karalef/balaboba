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

// MinTimeout is a minimum time limit for api requests.
const MinTimeout = 20 * time.Second

// ClientRus var.
var ClientRus = New(Rus)

// ClientEng var.
var ClientEng = New(Eng)

// New makes new balaboba api client.
//
// If the timeout is not specified or it is less than MinTimeout
// it will be equal to MinTimeout.
// Anyway the request can be canceled via the context.
func New(lang Lang, timeout ...time.Duration) *Client {
	d := net.Dialer{
		Timeout: MinTimeout,
	}
	if len(timeout) > 0 && timeout[0] > MinTimeout {
		d.Timeout = timeout[0]
	}
	return &Client{
		lang: lang,
		httpClient: http.Client{
			Timeout: d.Timeout,
			Transport: &http.Transport{
				DialTLSContext:      d.DialContext,
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

type responseBase struct {
	Error int `json:"error"`
}

func (r responseBase) err() int { return r.Error }

type errorable interface{ err() int }

func (c *Client) do(endpoint string, data interface{}, dst errorable) error {
	return c.doContext(context.Background(), endpoint, data, dst)
}

func (c *Client) doContext(ctx context.Context, endpoint string, data interface{}, dst errorable) error {
	err := c.request(ctx, apiurl+endpoint, data, dst)
	if err != nil {
		return err
	}
	if c := dst.err(); c != 0 {
		err = fmt.Errorf("balaboba: error code %d", c)
	}
	return err
}

func (c *Client) request(ctx context.Context, url string, data, dst interface{}) error {
	method := http.MethodGet
	var body io.Reader

	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return err
		}
		body = bytes.NewReader(b)
		method = http.MethodPost
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("balaboba: response status %s (%d)", resp.Status, resp.StatusCode)
	}

	if dst == nil {
		return nil
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(dst); err != nil {
		raw, err := io.ReadAll(io.MultiReader(dec.Buffered(), resp.Body))
		if err != nil {
			return err
		}
		return fmt.Errorf("response: %s, error: %w", string(raw), err)
	}
	return err
}

// Lang represents balaboba language.
type Lang uint8

// available languages.
const (
	Rus Lang = iota
	Eng
)
