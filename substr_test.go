package strtools

import (
	"testing"
)

func BenchmarkSubstr(b *testing.B) {
	const text = "Hello World"
	for n := 0; n < b.N; n++ {
		Substr(text, 4, len(text))
	}
}

type testpair struct {
	text     string
	start    int
	end      int
	expected string
}

var tests = []testpair{
	{"Hey hey", 0, 3, "Hey"},
	{"This is a -test", -4, 15, "test"},
	{"No text", 0, 0, ""},
	{"All the text", 0, 12, "All the text"},
}

func TestSubstr(t *testing.T) {
	for _, pair := range tests {
		result := Substr(pair.text, pair.start, pair.end)
		if result != pair.expected {
			t.Error(
				"For", pair.text,
				",",
				"expected", pair.expected,
				"got:", result,
			)
		}
	}
}
