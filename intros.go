package balaboba

import (
	"encoding/json"
)

// Intro is generation style.
type Intro struct {
	Style       uint8
	String      string
	Description string
}

// UnmarshalJSON is json.Unmarshaler interface implementation.
func (i *Intro) UnmarshalJSON(b []byte) error {
	var rep [3]interface{}
	err := json.Unmarshal(b, &rep)
	if err != nil {
		return err
	}

	i.Style = uint8(rep[0].(float64))
	i.String = rep[1].(string)
	i.Description = rep[2].(string)

	return nil
}

// IntrosResponse represents intros response.
type IntrosResponse struct {
	Intros []Intro `json:"intros"`
	Error  int     `json:"error"`
}

// Intros returns list of avaible generating styles.
func (c *Client) Intros() (*IntrosResponse, error) {
	ep := "intros"
	if c.lang == Eng {
		ep = "intros_eng"
	}
	var resp IntrosResponse
	return &resp, c.do(nil, ep, nil, &resp)
}
