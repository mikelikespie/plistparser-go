package plistparser

import (
	"testing"
)

func TestWrite(t *testing.T) {
	s, err := WriteString(map[string]interface{}{
		"foo": "bar",
		"baz": map[string]interface{}{
			".cheese": []interface{}{
				"cat\uFFFD",
				"3",
				"hello",
				map[string]interface{}{
					"hi":    "there",
					"wha t": "up",
					"z":     "",
				},
				[]byte("FOO"),
			},
		},
	})

	if err != nil {
		t.Error(err)
	}

	expected := `{
  baz = {
    .cheese = (
      "cat�",
      3,
      hello,
      {
        hi = there;
        "wha t" = up;
        z = "";
      },
      <46 4f 4f>,
    );
  };
  foo = bar;
}
`

	if s != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, s)
	}
}
