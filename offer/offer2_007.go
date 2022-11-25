package offer

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		target := 0 - nums[i]
		for l < r {
			var sum = nums[l] + nums[r]
			if sum == target {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r {
					l++
					if nums[l-1] != nums[l] {
						break
					}
				}
				for l < r {
					r--
					if nums[r] != nums[r+1] {
						break
					}
				}
			} else if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
