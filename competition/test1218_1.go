package competition

func SimilarPairs(words []string) int {
	hashmap := map[int]int{}
	for _, v := range words {
		num := 0
		for i := 0; i < len(v); i++ {
			num |= (1 << (v[i] - 'a'))
		}

		hashmap[num]++

	}
	res := 0
	for _, value := range hashmap {
		res += (value * (value - 1) / 2)
	}
	return res
}
