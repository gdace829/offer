package offer

import "sort"

func singleNonDuplicate(nums []int) int {
	length := len(nums)
	idx := sort.Search(length/2, func(i int) bool {
		return i*2 == length-1 || nums[i] != nums[i+1]
	})
	return nums[idx*2]
}
