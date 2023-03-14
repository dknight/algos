package main

import (
	"math/rand"
	"testing"
)

var qs []int

func init() {
	n := 100
	for i := 0; i < n; i++ {
		qs = append(qs, rand.Intn(n))
	}
}

func BenchmarkSimpleStocksSpan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SimpleStocksSpan(qs)
	}
}

func BenchmarkStackSlicesStocksSpan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StackSlicesStocksSpan(qs)
	}
}

func BenchmarkStackStocksSpan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StackStocksSpan(qs)
	}
}

func BenchmarkListStocksSpan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ListStocksSpan(qs)
	}
}
