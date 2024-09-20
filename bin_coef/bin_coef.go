package bin_coef

func BinCoef(r, c int) int {
	binCoefs := make([]int, 1, c+1)
	binCoefs[0] = 1
	for j := 1; j <= r; j++ {
		binCoefs = append(binCoefs, 1)
		for i := len(binCoefs) - 2; i > 0; i-- {
			binCoefs[i] = binCoefs[i-1] + binCoefs[i]
		}
	}
	return binCoefs[c]
}
