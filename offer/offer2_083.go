package offer

func permute(nums []int) [][]int {
	var dfs func()
	res := [][]int{}
	cur := []int{}
	vis := make([]bool, len(nums))
	dfs = func() {
		if len(cur) == len(nums) {
			res = append(res, append([]int{}, cur...))
		}
		for i := 0; i < len(nums); i++ {
			if vis[i] {
				continue
			}
			vis[i] = true
			cur = append(cur, nums[i])
			dfs()
			cur = cur[:len(cur)-1]
			vis[i] = false
		}

	}
	return res
}
