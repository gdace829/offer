package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	var find func(*TreeNode, int) bool
	find = func(tn *TreeNode, i int) bool {
		if tn == nil {
			return false
		}
		if tn.Val > i {
			return find(tn.Left, i)
		} else if tn.Val == i {
			return true
		} else {
			return find(tn.Right, i)
		}

	}
	var dfs func(*TreeNode) bool
	dfs = func(tn *TreeNode) bool {
		if tn.Left != nil {
			if dfs(tn.Left) {
				return true
			}
		}
		if 2*tn.Val < k {
			if find(root, k-tn.Val) {
				return true
			}
		} else {
			return false
		}
		if tn.Right != nil {
			if dfs(tn.Right) {
				return true
			}
		}
		return false
	}
	return dfs(root)
}
