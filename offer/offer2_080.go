package offer

func combine(n int, k int) [][]int {
	res := [][]int{}
	nums := []int{}
	var dfs func(int)
	dfs = func(now int) {
		if len(nums) == k {
			cur := make([]int, k)
			for i := 0; i < k; i++ {
				cur[i] = nums[i]
			}
			res = append(res, cur)
			return
		}

		for j := now; j <= n; j++ {
			nums = append(nums, j)
			dfs(j + 1)
			nums = nums[:len(nums)-1]
		}

	}
	dfs(1)
	return res
}
