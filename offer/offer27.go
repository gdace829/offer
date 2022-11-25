package offer

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newtree := new(TreeNode)
	newtree.Val = root.Val
	newtree.Left = mirrorTree(root.Right)
	newtree.Right = mirrorTree(root.Left)
	return newtree
}
