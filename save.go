package balaboba

import (
	"context"
)

// Saved contains link to the result.
type Saved struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

// SaveResult generates a link to the successful result.
func (c *Client) SaveResult(r Response) (*Saved, error) {
	z, err := c.zeliboba(r.raw.Query, r.raw.Text)
	if err != nil {
		return nil, err
	}
	var s struct {
		responseBase
		Saved
	}
	err = c.do("save2", struct {
		response
		ID   string `json:"id"`
		Mode string `json:"mode"`
	}{r.raw, z.DefaultEn, "DONE"}, &s)
	if err != nil {
		return nil, err
	}

	return &s.Saved, nil
}

type zelibobaResponse struct {
	Default   string `json:"default"`
	DefaultEn string `json:"default-en"`
	VK        string `json:"vk"`
	VKEn      string `json:"vk-en"`
}

func (c *Client) zeliboba(q, text string) (*zelibobaResponse, error) {
	const zelibobaURL = "https://yandex.ru/lab/research-pic-generator/zeliboba/"

	var r zelibobaResponse
	err := c.request(context.Background(), zelibobaURL, [2]string{q, text}, &r)
	if err != nil {
		return nil, err
	}
	return &r, err
}
