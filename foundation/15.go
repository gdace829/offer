package foundation

import "sort"

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	ans := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if nums[k]+nums[j]+nums[i] == 0 {
				ans = append(ans, []int{nums[j], nums[i], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[k]+nums[j]+nums[i] < 0 {
				j++
			} else if nums[k]+nums[j]+nums[i] > 0 {
				k--
			}

		}

	}
	return ans
}
