package offer

import "math"

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	var max int
	leaves := []*TreeNode{}
	leaves = append(leaves, root)

	for len(leaves) != 0 {
		max = math.MinInt32
		leaves2 := []*TreeNode{}
		for len(leaves) != 0 {
			if leaves[0].Val > max {
				max = leaves[0].Val
			}
			if leaves[0].Left != nil {
				leaves2 = append(leaves2, leaves[0].Left)
			}
			if leaves[0].Right != nil {
				leaves2 = append(leaves2, leaves[0].Right)
			}
			leaves = leaves[1:]
		}
		res = append(res, max)
		leaves = leaves2
	}
	return res
}
