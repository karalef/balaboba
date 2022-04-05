package balaboba

import "context"

const text3api = "text3"

// GetResponse represents text generating response.
type GetResponse struct {
	Query    string `json:"query"`
	Text     string `json:"text"`
	BadQuery int    `json:"bad_query"`
	Error    int    `json:"error"`
}

// Get returns text gerated with passed params.
func (c *Client) Get(ctx context.Context, query string, style Style, filter ...int) (*GetResponse, error) {
	get := struct {
		Query  string `json:"query"`
		Style  Style  `json:"intro"`
		Filter int    `json:"filter"`
	}{
		Query:  query,
		Style:  style,
		Filter: 1,
	}
	if len(filter) > 0 {
		get.Filter = filter[0]
	}
	var resp GetResponse
	return &resp, c.do(ctx, text3api, get, &resp)
}
