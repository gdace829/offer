package offer

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
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

		nums = append(nums, candidates[index])
		dfs(index+1, ramin-candidates[index])
		nums = nums[:len(nums)-1]
		for index < len(candidates)-1 {
			if candidates[index] == candidates[index+1] {
				index++
			} else {
				break
			}
		}
		dfs(index+1, ramin)
	}
	dfs(0, target)
	return res
}
