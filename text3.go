package balaboba

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

const text3api = "text3"

// GetRequest contains parameters for text generating.
type GetRequest struct {
	Query  string `json:"query"`
	Filter int    `json:"filter"`
	Style  Style  `json:"intro"`
}

// GetResponse represents text generating response.
type GetResponse struct {
	Query    string `json:"query"`
	Text     string `json:"text"`
	BadQuery int    `json:"bad_query"`
	Error    int    `json:"error"`
}

// Get returns text gerated with passed params.
func (c *Client) Get(get GetRequest) (*GetResponse, error) {
	p, err := json.Marshal(get)
	if err != nil {
		return nil, err
	}

	res, err := c.c.Post(api+text3api, "application/json", bytes.NewReader(p))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resp GetResponse
	return &resp, json.NewDecoder(res.Body).Decode(&resp)
}

// Options ???
func (c *Client) Options() error {
	req, err := http.NewRequest(http.MethodOptions, api+text3api, nil)
	if err != nil {
		return err
	}
	res, err := c.c.Do(req)
	if err != nil {
		return err
	}
	res.Body.Close()

	if res.StatusCode != http.StatusOK {
		err = errors.New(res.Status)
	}
	return err
}
