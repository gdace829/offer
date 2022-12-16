package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	var dfs func(*TreeNode, int)
	res := 0
	dfs = func(tn *TreeNode, i int) {
		if tn.Left == nil && tn.Right == nil {
			res += tn.Val + i*10
			return
		}
		if tn.Left != nil {
			dfs(tn.Left, i*10+tn.Val)
		}
		if tn.Right != nil {
			dfs(tn.Right, i*10+tn.Val)
		}
	}
	dfs(root, 0)
	return res
}
