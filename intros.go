package balaboba

import (
	"encoding/json"
)

// Intro is a generation style.
type Intro struct {
	Style       int
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

	i.Style = int(rep[0].(float64))
	i.String = rep[1].(string)
	i.Description = rep[2].(string)

	return nil
}

// Intros returns list of avaible generating styles.
func (c *Client) Intros() ([]Intro, error) {
	ep := "intros"
	if c.lang == Eng {
		ep = "intros_eng"
	}
	var resp struct {
		responseBase
		Intros []Intro `json:"intros"`
	}
	return resp.Intros, c.do(ep, nil, &resp)
}
