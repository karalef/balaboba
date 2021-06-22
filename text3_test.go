package balaboba

import "testing"

func TestGet(t *testing.T) {
	c := New()

	gen, err := c.Get(GetRequest{
		Query:  "123",
		Filter: 1,
		Style:  NoStyle,
	})
	if err != nil {
		t.Fatal(err)
	}

	if gen.BadQuery != 0 {
		t.Log("bad query", gen.BadQuery)
		t.FailNow()
	}

	if gen.Error != 0 {
		t.Log("bad error", gen.Error)
		t.FailNow()
	}

	if gen.Text == "" {
		t.Fatal()
	}

	t.Log(gen.Text)
}

func TestOptions(t *testing.T) {
	c := New()

	err := c.Options()
	if err != nil {
		t.Fatal(err)
	}
}
