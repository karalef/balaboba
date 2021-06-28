package balaboba

import (
	"encoding/json"
)

const introsapi = "intros"

// Intro is generating style.
type Intro struct {
	Style       Style
	String      string
	Description string
}

// UnmarshalJSON is Unmarshaler interface implementation.
func (i *Intro) UnmarshalJSON(b []byte) error {
	var rep [3]interface{}
	err := json.Unmarshal(b, &rep)
	if err != nil {
		return err
	}

	i.Style = Style(rep[0].(float64))
	i.String = rep[1].(string)
	i.Description = rep[2].(string)

	return nil
}

// IntrosResponse represents intros reponse.
type IntrosResponse struct {
	Intros []Intro `json:"intros"`
	Error  int     `json:"error"`
}

// Intros returns list of avaible generating styles.
func (c *Client) Intros() (*IntrosResponse, error) {
	var resp IntrosResponse
	return &resp, c.do(introsapi, nil, &resp)
}
