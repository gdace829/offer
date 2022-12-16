package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var dfs func(*TreeNode)
	isfind := false
	isres := false
	var res *TreeNode
	dfs = func(tn1 *TreeNode) {
		if tn1 == nil {
			return
		}
		dfs(tn1.Left)
		if isfind == true && isres == false {
			res = tn1
			isres = true
			return
		}
		if tn1 == p {
			isfind = true
		}
		dfs(tn1.Right)

	}
	dfs(root)
	return res
}
