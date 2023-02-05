package offer

import "sort"

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	sort.Ints(nums)
	for k, _ := range dp {
		for _, v := range nums {
			if k >= v {
				dp[k] = dp[k-v] + dp[k]
			} else {
				break
			}
		}
	}
	return dp[target]
}
