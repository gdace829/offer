package offer

import (
	"sort"
)

func minEatingSpeed(piles []int, h int) int {
	res := sort.Search(int(1e9), func(i int) bool {
		if i == 0 {
			return false
		}
		var total int
		for _, v := range piles {
			total += (v + i - 1) / i
		}
		return total <= h
	})
	return res
}
