func letterCombinations(digits string) []string {
	keys := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}

	D := len(digits)
	possibilities := make([]int, D)
	possibilitiesSoFar := make([]int, D)
	N := 0

	for i := 0; i < len(digits); i++ {
		possibilities[i] = len(keys[digits[i]])
		possibilitiesSoFar[i] = len(keys[digits[i]])
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
			resI[d] = keys[digits[d]][(i/(N/possibilitiesSoFar[d]))%len(keys[digits[d]])]
		}

		res[i] = string(resI)
	}

	return res
}