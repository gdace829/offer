package offer

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if A.Val == B.Val && helper(A, B) {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func helper(A *TreeNode, B *TreeNode) bool {
	if A == nil && B != nil {
		return false
	}
	if B == nil {
		return true
	}
	if A.Val == B.Val {
		return helper(A.Left, B.Left) && helper(A.Right, B.Right)
	}
	return false
}
