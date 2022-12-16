package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	res := []int{}
	for len(stack) != 0 {
		stack1 := []*TreeNode{}
		for i := 0; i < len(stack); i++ {
			if stack[i].Left != nil {
				stack1 = append(stack1, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack1 = append(stack1, stack[i].Right)
			}
		}
		res = append(res, stack[len(stack)-1].Val)
		stack = stack1
	}

	return res
}
