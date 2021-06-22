package balaboba

import "testing"

func TestIntros(t *testing.T) {
	c := New()

	intros, err := c.Intros()
	if err != nil {
		t.Fatal(err)
	}

	if intros.Error != 0 {
		t.Log("reponse error", intros.Error)
		t.FailNow()
	}

	for _, i := range intros.Intros {
		t.Log(i.String(), i.Description(), i.Style())
	}
}
