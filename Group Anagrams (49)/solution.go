func groupAnagrams(strs []string) [][]string {
    wordsMap := map[string][]string{}

    for _, word := range strs {
        normalizedArr := strings.Split(word, "")
        sort.Strings(normalizedArr)
        normalized := strings.Join(normalizedArr, "")

        anagrams := wordsMap[normalized]
        anagrams = append(anagrams, word)
        wordsMap[normalized] = anagrams
    }

    res := [][]string{}
    for _, list := range wordsMap {
        res = append(res, list)
    }

    return res
}