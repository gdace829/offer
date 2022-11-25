package offer

func findMaxLength(nums []int) int {
	hashmap := make(map[int]int)
	hashmap[0] = -1
	onemore := 0
	res := 0
	for index, num := range nums {
		if num == 1 {
			onemore++
		} else {
			onemore--
		}
		if numindex, is := hashmap[onemore]; is {
			res = max(res, index-numindex)
		} else {
			hashmap[onemore] = index
		}

	}
	return res
}
