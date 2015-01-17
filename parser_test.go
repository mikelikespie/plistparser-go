package plistparser

import (
	"testing"
)

func TestParse(t *testing.T) {
	r, err := ParseString(`{abc=cde;}`)
	if err != nil {
		t.Errorf("Should have parsed, but got error %s", err)
	}
	switch r := r.(type) {
	case map[string]interface{}:
		if r["abc"] != "cde" {
			t.Errorf("expected key abc to equal cde ")
		}
	default:
		t.Errorf("unexpected root type")
	}
}

func TestSlashes(t *testing.T) {
	r, err := ParseString(`/ab/c`)
	if err != nil {
		t.Fatalf("Should have parsed, but got error %s", err)
	}
	switch r := r.(type) {
	case string:
		if r != `/ab/c` {
			t.Errorf("expected valid string parse")
		}
	default:
		t.Errorf("unexpected root type")
	}
}

func TestArray1(t *testing.T) {
	r, err := ParseString(`(abc)`)
	if err != nil {
		t.Fatalf("Should have parsed, but got error %s", err)
	}
	switch r := r.(type) {
	case []interface{}:
		if r[0] != `abc` {
			t.Errorf("expected valid string parse")
		}
	default:
		t.Errorf("unexpected root type")
	}
}

func TestArray2(t *testing.T) {
	r, err := ParseString(`(abc,)`)
	if err != nil {
		t.Fatalf("Should have parsed, but got error %s", err)
	}
	switch r := r.(type) {
	case []interface{}:
		if r[0] != `abc` {
			t.Errorf("expected valid string parse")
		}
	default:
		t.Errorf("unexpected root type")
	}
}

func TestErrors(t *testing.T) {
	_, err := ParseString(`{abc=cde}`)
	if err == nil {
		t.Errorf("expected error")
	}

	_, err = ParseString(`{abc=(,);}`)
	if err == nil {
		t.Errorf("expected error")
	}
	_, err = ParseString(`{abc=<aaaa bbbg>;}`)
	if err == nil {
		t.Errorf("expected error")
	}

	_, err = ParseString(`{abc b b=<aaaa bbbb>;}`)
	if err == nil {
		t.Errorf("expected error")
	}

	_, err = ParseString(`(a b)`)
	if err == nil {
		t.Errorf("expected error")
	}

	_, err = ParseString(`(a,, b)`)
	if err == nil {
		t.Errorf("expected error")
	}
	_, err = ParseString(`(a, b,,)`)
	if err == nil {
		t.Errorf("expected error")
	}
}
