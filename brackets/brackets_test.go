package brackets

import "testing"

var inputs = map[string]bool{
	"":         false,
	"()":       true,
	"[]":       true,
	"{}":       true,
	"([])":     true,
	"([]":      false,
	"([](){})": true,
	"[])}{}()": false,
	"{([])}":   true,
	"{)":       false,
}

func TestAreBalanced(t *testing.T) {
	for input, exp := range inputs {
		got := AreBalanced(input)
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	}
}
