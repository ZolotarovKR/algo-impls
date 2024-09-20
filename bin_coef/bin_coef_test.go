package bin_coef

import "testing"

func TestBinCoef(t *testing.T) {
	testCases := []struct{ r, c, expected int }{
		{6, 3, 20},
		{0, 0, 1},
		{4, 4, 1},
		{5, 1, 5},
		{3, 2, 3},
		{6, 2, 15},
	}
	for _, testCase := range testCases {
		if got := BinCoef(testCase.r, testCase.c); got != testCase.expected {
			t.Fatalf(
				"Expected BinCoef(%v, %v) to equal %v, but got %v",
				testCase.r,
				testCase.c,
				testCase.expected,
				got,
			)
		}
	}

}
