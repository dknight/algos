package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dknight/algo-ads/stack"
)

type PostfixNotation string

func (pn PostfixNotation) Parse() []string {
	ss := strings.Fields(string(pn))
	return ss
}

func (pn PostfixNotation) Calculate() float64 {
	result := 0.0
	stk := stack.New[float64]()
	for _, s := range pn.Parse() {
		x, err := strconv.ParseFloat(s, 32)
		if err != nil {
			if !isOperand(s) {
				panic("Parse error")
			}
			a := stk.Pop()
			b := stk.Pop()
			// fmt.Println(x, s, y)
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
	return (s == "-" || s == "+" || s == "/" || s == "*")
}

var (
	input = "1 2 3 * + 2 - 10 3 4 / 5 + 1 +"
)

func main() {
	s := PostfixNotation(input)
	result := s.Calculate()
	fmt.Println(result)
}
