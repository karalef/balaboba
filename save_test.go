package balaboba

import "testing"

func TestSave(t *testing.T) {
	c := ClientRus

	gen, err := c.Generate("123", Standart)
	if err != nil {
		t.Fatal(err)
	}
	if gen.BadQuery {
		t.Fatal("bad query", gen.BadQuery)
	}

	_, err = c.SaveResult(*gen)
	if err != nil {
		t.Fatal(err)
	}
}
