package offer

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	hashmap := make(map[int]int)
	for key, v := range arr2 {
		hashmap[v] = key
	}
	sort.Slice(arr1, func(i, j int) bool {
		a, err := hashmap[arr1[i]]
		b, err1 := hashmap[arr1[j]]
		if err == false && err1 == false {
			return arr1[i] < arr1[j]
		}
		if err == false {
			return false
		}
		if err1 == false {
			return true
		}
		return a < b
	})
	return arr1
}
