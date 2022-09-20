package balaboba

import "testing"

func TestIntros(t *testing.T) {
	c := ClientRus

	intros, err := c.Intros()
	if err != nil {
		t.Fatal(err)
	}
	if len(intros) == 0 {
		t.Fatal("response does not contain intros")
	}
}
