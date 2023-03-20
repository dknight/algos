// Package factorial represents examples of calculations of factorial (n!).
// Read more about factorial: https://en.wikipedia.org/wiki/Factorial
package factorial

import "golang.org/x/exp/constraints"

// Recursive calculates the factorial in unsigned way.
func Recursive[T constraints.Integer](x T) T {
	if x < 0 {
		panic("Negative value in factorial")
	}
	if x == 0 || x == 1 {
		return 1
	}
	return x * Recursive(x-1)
}

// Calculate calculates factorial in imperative way. This way is much more
// performing than recursive way.
func Calculate[T constraints.Integer](x T) T {
	if x < 0 {
		panic("Negative value in factorial")
	}
	var i, cur, prev T
	cur, prev = 1, 1
	for i = 1; i <= x; i++ {
		cur = prev * i
		prev = cur
	}
	return cur
}
