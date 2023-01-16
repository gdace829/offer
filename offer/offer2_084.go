package offer

import "sort"

func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}
	vis := make([]bool, len(nums))
	sort.Ints(nums)
	var dfs func()
	dfs = func() {
		if len(cur) == len(nums) {
			res = append(res, append([]int{}, cur...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if vis[i] {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !vis[i-1] {
				continue
			}
			vis[i] = true
			cur = append(cur, nums[i])
			dfs()
			vis[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	dfs()
	return res
}
