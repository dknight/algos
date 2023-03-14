package stocks

import (
	"math/rand"
	"reflect"
	"testing"
)

var benchQs []int

func init() {
	n := 10000
	for i := 0; i < n; i++ {
		benchQs = append(benchQs, rand.Intn(n))
	}
}

func BenchmarkStocksSpansSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SpansSlices(benchQs)
	}
}

func BenchmarkStocksSpansStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SpansStack(benchQs)
	}
}

type TestPair struct {
	Quotes   []int
	Expected []int
}

func TestBenchmarkStocksSpans(t *testing.T) {
	pairs := []TestPair{
		TestPair{
			[]int{7, 11, 8, 6, 3, 8, 9},
			[]int{1, 2, 1, 1, 1, 4, 5},
		},
		TestPair{
			[]int{5, 3, 2, 3},
			[]int{1, 1, 1, 3},
		},
		TestPair{
			[]int{0, 1, 2, 3},
			[]int{1, 2, 3, 4},
		},
		TestPair{
			[]int{3, 1, 2, 3},
			[]int{1, 1, 2, 4},
		},
		// TODO more
	}
	t.Run("SpansSlices", func(t *testing.T) {
		for _, pair := range pairs {
			got := SpansSlices(pair.Quotes)
			if !reflect.DeepEqual(got, pair.Expected) {
				t.Errorf("Expected %v got %v", pair.Expected, got)
			}
		}
	})
	t.Run("SpansStack", func(t *testing.T) {
		for _, pair := range pairs {
			got := SpansStack(pair.Quotes)
			if !reflect.DeepEqual(got, pair.Expected) {
				t.Errorf("Expected %v got %v", pair.Expected, got)
			}
		}
	})
}
