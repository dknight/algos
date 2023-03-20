package factorial

import "testing"

func assertPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if err := recover(); err == nil {
			t.Error(err)
		}
	}()
	fn()
}

func TestRecursive(t *testing.T) {
	t.Run("Zero", func(t *testing.T) {
		exp := 1
		got := Recursive[int](0)
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("One", func(t *testing.T) {
		exp := 1
		got := Recursive[int](1)
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("Common", func(t *testing.T) {
		exp := 120
		got := Recursive[int](5)
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}

		exp = 3628800
		got = Recursive[int](10)
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("Panic", func(t *testing.T) {
		assertPanic(t, func() {
			Recursive(-42)
		})
	})
}

func TestCalculate(t *testing.T) {
	t.Run("Zero", func(t *testing.T) {
		exp := 1
		got := Calculate[int](0)
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("One", func(t *testing.T) {
		exp := 1
		got := Calculate[int](1)
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("Common", func(t *testing.T) {
		exp := 120
		got := Calculate[int](5)
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}

		exp = 3628800
		got = Calculate[int](10)
		if exp != got {
			t.Errorf("Expected %v got %v", exp, got)
		}
	})

	t.Run("Panic", func(t *testing.T) {
		assertPanic(t, func() {
			Calculate(-42)
		})
	})
}

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recursive(10)
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(10)
	}
}
