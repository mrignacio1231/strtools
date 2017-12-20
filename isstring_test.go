package strtools

import "testing"

type testpair2 struct {
	text   interface{}
	result bool
}

var tests2 = []testpair2{
	{"Hello", true},
	{1.2, false},
	{5, false},
	{'f', false},
}

func TestIsString(t *testing.T) {

	for _, pair := range tests2 {
		if IsString(pair.text) != pair.result {
			t.Error(pair.text, "!=", pair.result)
		}
	}
}
