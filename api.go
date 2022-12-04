package balaboba

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const apiurl = "https://yandex.ru/lab/api/yalm/"

// ClientRus is default russian client.
var ClientRus = NewWithTimeout(Rus, 20*time.Second)

// ClientEng is default english client.
var ClientEng = NewWithTimeout(Eng, 20*time.Second)

// New makes new balaboba api client.
func New(lang Lang, client ...*http.Client) *Client {
	c := &Client{
		HTTP: http.DefaultClient,
		lang: lang,
	}
	if len(client) > 0 && client[0] != nil {
		c.HTTP = client[0]
	}
	return c
}

// NewWithTimeout makes new balaboba api client with resuests timeout.
func NewWithTimeout(lang Lang, timeout time.Duration) *Client {
	return New(lang, &http.Client{Timeout: timeout})
}

// Client is Yandex Balaboba client.
type Client struct {
	HTTP *http.Client
	lang Lang
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
		var w *io.PipeWriter
		body, w = io.Pipe()
		go func() { w.CloseWithError(json.NewEncoder(w).Encode(data)) }()
		method = http.MethodPost
	}

	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("balaboba: response status %s", resp.Status)
	}

	if dst == nil {
		return nil
	}

	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(dst); err != nil {
		raw, _ := io.ReadAll(io.MultiReader(dec.Buffered(), resp.Body))
		err = fmt.Errorf("balaboba: %s\nresponse: %s", err.Error(), string(raw))
	}
	return err
}

// Lang returns client language.
func (c Client) Lang() Lang {
	return c.lang
}

// Warn1 returns a first warning in the client language.
func (c Client) Warn1() string {
	if c.lang == Rus {
		return Warn1Rus
	}
	return Warn1Eng
}

// Warn2 returns a second warning in the client language.
func (c Client) Warn2() string {
	if c.lang == Rus {
		return Warn2Rus
	}
	return Warn2Eng
}

// About returns a text about Balaboba in the client language.
func (c Client) About() string {
	if c.lang == Rus {
		return AboutRus
	}
	return AboutEng
}

// BadQuery returns a bad query response in the client language.
func (c Client) BadQuery() string {
	if c.lang == Rus {
		return BadQueryRus
	}
	return BadQueryEng
}

// Lang represents balaboba language.
type Lang uint8

// available languages.
const (
	Rus Lang = iota
	Eng
)
