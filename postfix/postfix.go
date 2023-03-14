package postfix

import (
	"strconv"
	"strings"

	"github.com/dknight/algos/stack"
)

// Notation represents postfix notation string.
type Notation string

// Parse postfix notation expression into numbers and operands.
func (pn Notation) Parse() []string {
	return strings.Fields(string(pn))
}

// Calculate the expression of the postfix notations.
func (pn Notation) Calculate() float64 {
	result := 0.0
	stk := stack.New[float64]()
	for _, s := range pn.Parse() {
		x, err := strconv.ParseFloat(s, 64)
		if err != nil {
			if !isOperand(s) {
				panic("Parse error")
			}
			a := stk.Pop()
			b := stk.Pop()
			switch s {
			case "-":
				result = b - a
			case "+":
				result = b + a
			case "*":
				result = b * a
			case "/":
				if a == 0 {
					panic("Division by zero")
				}
				result = b / a
			}
			stk.Push(result)
		} else {
			stk.Push(x)
		}
	}
	return result
}

func isOperand(s string) bool {
	switch s {
	case "-", "+", "/", "*":
		return true
	default:
		return false
	}
}
