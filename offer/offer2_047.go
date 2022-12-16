package offer

func pruneTree(root *TreeNode) *TreeNode {

	var leavesok func(*TreeNode) bool
	leavesok = func(tn *TreeNode) bool {
		if tn == nil {
			return false
		}
		isleft := leavesok(tn.Left)
		if isleft == false {
			tn.Left = nil

		}
		isright := leavesok(tn.Right)
		if isright == false {
			tn.Right = nil

		}
		return (tn.Val == 1) || isright || isleft
	}
	if leavesok(root) == false {
		return nil
	} else {
		return root
	}

}
