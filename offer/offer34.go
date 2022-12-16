package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res [][]int

func pathSum1(root *TreeNode, target int) [][]int {
	res = make([][]int, 0)
	path(root, target, 0, make([]int, 0))
	return res
}

func path(root *TreeNode, target, sum int, res1 []int) {
	if root == nil {
		return
	}
	if sum+root.Val == target && root.Left == nil && root.Right == nil {
		res1 = append(res1, root.Val)
		res0 := make([]int, len(res1))
		copy(res0, res1)
		res = append(res, res0)
		res1 = res1[:len(res1)-1]
		return
	}
	res1 = append(res1, root.Val)
	path(root.Right, target, sum+root.Val, res1)
	path(root.Left, target, sum+root.Val, res1)
	res1 = res1[:len(res1)-1]
}
