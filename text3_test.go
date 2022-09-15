package balaboba

import "testing"

func TestGenerateNormal(t *testing.T) {
	c := ClientRus

	// normal request
	gen, err := c.Generate(nil, "123", Standart)
	if err != nil {
		t.Fatal(err)
	}
	if gen.Error != 0 {
		t.Fatal("bad error", gen.Error)
	}
	if gen.BadQuery != 0 {
		t.Fatal("bad query", gen.BadQuery)
	}
}
