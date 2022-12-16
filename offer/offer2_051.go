package offer

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	var dfs func(*TreeNode) int
	var max func(int, int) int
	res := math.MinInt32
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		} else {
			return i2
		}
	}
	dfs = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		l := dfs(tn.Left)
		r := dfs(tn.Right)
		res = max(res, max(tn.Val+l+r, max(tn.Val+l, tn.Val+r)))
		res = max(res, tn.Val)
		return max(tn.Val, max(tn.Val+l, tn.Val+r))

	}
	dfs(root)
	return res

}
