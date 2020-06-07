package test

import (
	"fmt"
	"testing"
)

func TestIntMin(t *testing.T) {
	var tests = []struct {
		a, b,
		expected int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for index, tt := range tests {
		// testName := fmt.Sprintf("%d, %d", tt.a, tt.b)
		testName := fmt.Sprintf("IntMin#%d", index)
		t.Run(testName, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.expected {
				t.Errorf("got %d, expected %d", ans, tt.expected)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	num1, num2 := -1, 0
	expected := -1

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		actual := IntMin(num1, num2)
		if actual != expected {
			b.Errorf("got %d, expected %d", actual, expected)
		}
	}
}
