package offer

func isSymmetric(root *TreeNode) bool {
	return isSame(root.Right, root.Left)
}

func isSame(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}
	if a.Val == b.Val {
		return isSame(a.Right, b.Left) && isSame(a.Left, b.Right)
	} else {
		return false
	}
}
