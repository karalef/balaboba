package balaboba

import "testing"

func TestIntros(t *testing.T) {
	c := ClientRus

	intros, err := c.Intros()
	if err != nil {
		t.Fatal(err)
	}

	if intros.Error != 0 {
		t.Fatal("reponse error", intros.Error)
	}
}
