package brackets

import (
	"strings"

	"github.com/dknight/algos/structures/stack"
)

var openCloseMap = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

// AreBalanced checks that brackets are balanced or not.
func AreBalanced(input string) bool {
	if len(input) == 0 {
		return false
	}
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
			for o, c := range openCloseMap {
				if last != o && s == c {
					return false
				}
			}
		}
	}
	return stk.Len() == 0
}
