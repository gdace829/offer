package offer

import "math"

func minSubArrayLen(target int, nums []int) int {
	r := -1
	l := 0
	res := math.MaxInt32
	cur := 0
	for r < len(nums)-1 {
		r++
		cur += nums[r]
		for cur >= target {
			if r-l+1 < res {

				res = r - l + 1
			}
			cur -= nums[l]
			l++
		}
	}
	if res == math.MaxInt32 {
		return 0
	} else {
		return res
	}
}
