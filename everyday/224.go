package everyday

import (
	"math"
)

// 2357 使数组中所有元素都等于零
func minimumOperations(nums []int) int {
	min := math.MaxInt
	res := 0
	count := len(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 && nums[i] < min {
			min = nums[i]
		}
		if nums[i] == 0 {
			count--
		}
	}
	if count == 1 && nums[0] == 0 {
		return 0
	}
	for count != 0 {
		min2 := math.MaxInt
		count = len(nums)
		for i := 0; i < len(nums); i++ {
			if nums[i] != 0 {
				nums[i] = nums[i] - min
			}
			if nums[i] != 0 {
				if nums[i] < min2 {
					min2 = nums[i]
				}
			} else {
				count--
			}
		}
		min = min2
		res++
	}
	return res

}
