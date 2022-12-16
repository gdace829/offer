package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode)
	cur := 0
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		dfs(tn.Right)
		cur += tn.Val
		tn.Val = cur
		dfs(tn.Left)
	}
	dfs(root)
	return root
}
