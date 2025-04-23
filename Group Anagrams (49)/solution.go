func groupAnagrams(strs []string) [][]string {
	wordsMap := map[string]int{}
    res := [][]string{}
	cnt := 0
	for _, word := range strs {
		normalizedArr := []byte(word)
		sort.Slice(normalizedArr, func(i, j int) bool {
			return normalizedArr[i] < normalizedArr[j]
		})
		normalized := string(normalizedArr)

		id, ok := wordsMap[normalized]
        if !ok {
            id = cnt
            wordsMap[normalized] = id
            res = append(res, []string{})

			cnt++
		}

		res[id] = append(res[id], word)
	}

	return res
}