package main

import (
	"fmt"
	"strings"

	"github.com/dknight/algo-ads/stack"
)

var (
	input    = "{()([]{})}{}()()"
	mappings = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
)

func main() {
	result := BracketsProblemResolver(input)
	fmt.Println(result)
}

func BracketsProblemResolver(input string) bool {
	stk := stack.New[string]()
	ss := strings.Split(input, "")
	for _, s := range ss {
		if s == "(" || s == "[" || s == "{" {
			stk.Push(s)
			continue
		}
		if stk.Len() == 0 {
			return false
		}
		if s == ")" || s == "]" || s == "}" {
			last := stk.Pop()
			for o, c := range mappings {
				if last != o && s == c {
					return false
				}
			}
		}
	}
	return (stk.Len() == 0)
}
