package offer

func findBottomLeftValue1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var depfind func(*TreeNode) []int
	depfind = func(tn *TreeNode) []int {
		var a, b []int
		if tn == nil {
			return []int{0, -1}
		}
		if tn.Left == nil && tn.Right == nil {
			return []int{1, tn.Val}
		}
		a = depfind(tn.Left)

		b = depfind(tn.Right)

		if a[0] >= b[0] {
			return []int{a[0] + 1, a[1]}
		} else {
			return []int{b[0] + 1, b[1]}
		}

	}
	return depfind(root)[1]

}
func findBottomLeftValue(root *TreeNode) int {
	height := 0
	res := 0
	var dfs func(*TreeNode, int)
	dfs = func(tn *TreeNode, i int) {
		if tn == nil {
			return
		}
		if tn.Left == nil && tn.Right == nil {
			if i > height {
				height = i
				res = tn.Val
			}
		}

		dfs(tn.Left, i+1)
		dfs(tn.Right, i+1)

	}
	dfs(root, 1)
	return res

}
