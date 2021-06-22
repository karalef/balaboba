package balaboba

const text3api = "text3"

// GetRequest contains parameters for text generating.
type GetRequest struct {
	Query  string `json:"query"`
	Style  Style  `json:"intro"`
	Filter int    `json:"filter"`
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
	var resp GetResponse
	return &resp, c.post(text3api, &resp, get)
}

// Options ???
func (c *Client) Options() error {
	return c.options(text3api)
}
