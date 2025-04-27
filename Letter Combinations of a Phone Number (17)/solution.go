func letterCombinations(digits string) []string {
	keys := [][]byte{
		[]byte{'a', 'b', 'c'},
		[]byte{'d', 'e', 'f'},
		[]byte{'g', 'h', 'i'},
		[]byte{'j', 'k', 'l'},
		[]byte{'m', 'n', 'o'},
		[]byte{'p', 'q', 'r', 's'},
		[]byte{'t', 'u', 'v'},
		[]byte{'w', 'x', 'y', 'z'},
	}

	D := len(digits)
	possibilitiesSoFar := make([]int, D)
	N := 0

	for i := 0; i < len(digits); i++ {
		possibilitiesSoFar[i] = len(keys[digits[i]-'2'])
		N = possibilitiesSoFar[i]
	}

	for i := 1; i < len(possibilitiesSoFar); i++ {
		possibilitiesSoFar[i] *= possibilitiesSoFar[i-1]
		N = possibilitiesSoFar[i]
	}

	res := make([]string, N)

	resI := make([]byte, D)
	for i := 0; i < len(res); i++ {
		for d := 0; d < D; d++ {
			resI[d] = keys[digits[d]-'2'][(i/(N/possibilitiesSoFar[d]))%len(keys[digits[d]-'2'])]
		}

		res[i] = string(resI)
	}

	return res
}