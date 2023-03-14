package postfix

import "testing"

var inputs = map[string]float64{
	"2 2 +":                          4,
	"1 2 3 * + 2 - 10 3 4 / 5 + 1 +": 6.75,
	"0":                              0,
	"1 2 3 + -":                      -4,
	"6 2 / ":                         3,
	// TODO make more
}

func TestPostfix(t *testing.T) {
	for input, exp := range inputs {
		pn := Notation(input)
		got := pn.Calculate()
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	}
}
