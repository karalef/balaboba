package balaboba

import "encoding/json"

const introsapi = "intros"

// Intro is generating style.
type Intro [3]interface{}

// Style returns style code.
func (i Intro) Style() Style {
	return Style(i[0].(int))
}

func (i Intro) String() string {
	return i[1].(string)
}

// Description of style.
func (i Intro) Description() string {
	return i[2].(string)
}

// IntrosResponse represents intros reponse.
type IntrosResponse struct {
	Intros []Intro `json:"intros"`
	Error  int     `json:"error"`
}

// Intros returns list of avaible generating styles.
func (c *Client) Intros() (*IntrosResponse, error) {
	res, err := c.c.Get(api + introsapi)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resp IntrosResponse
	return &resp, json.NewDecoder(res.Body).Decode(&resp)
}
