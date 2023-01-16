package offer

func subsets(nums []int) [][]int {
	var dfs func([]int, int)
	res := [][]int{}
	res = append(res, []int{})
	dfs = func(cur []int, j int) {

		for i := j; i < len(nums); i++ {
			cur = append(cur, nums[i])
			cur1 := make([]int, len(cur))
			for j := 0; j < len(cur); j++ {
				cur1[j] = cur[j]
			}
			res = append(res, cur1)
			dfs(cur, i+1)
			cur = cur[:len(cur)-1]
		}
	}
	dfs([]int{}, 0)
	return res
}
