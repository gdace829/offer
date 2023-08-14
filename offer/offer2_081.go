package offer

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var dfs func(int, int)
	sort.Ints(candidates)
	res := [][]int{}
	nums := []int{}
	dfs = func(index int, ramin int) {
		if ramin == 0 {
			res = append(res, append([]int{}, nums...))
			return
		}
		if index == len(candidates) {
			return
		}
		if candidates[index] > ramin {
			return
		}
		dfs(index+1, ramin)
		nums = append(nums, candidates[index])
		dfs(index, ramin-candidates[index])
		nums = nums[:len(nums)-1]
	}
	return res
}
