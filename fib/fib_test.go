package fib

import "testing"

func TestFib(t *testing.T) {
	testCases := []struct{ n, expected int }{
		{10, 55},
		{40, 102334155},
		{1, 1},
		{20, 6765},
		{12, 144},
		{46, 1836311903},
		{0, 0},
	}
	for _, testCase := range testCases {
		if got := Fib(testCase.n); got != testCase.expected {
			t.Fatalf("Expected Fib(%v) to equal %v, but got %v", testCase.n, testCase.expected, got)
		}
	}
}
