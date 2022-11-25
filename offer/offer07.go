package offer

// 每一层为进入下一层做好铺垫比如计算出左右子树前序遍历中序遍历
func BuildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}
	var root int
	var leftinorder []int
	var rightinorder []int
	var leftpreorder []int
	var rightpreorder []int
	root = preorder[0]
	tree := new(TreeNode)
	tree.Val = root
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root {
			leftinorder = inorder[:i]
			rightinorder = inorder[i+1:]
			leftpreorder = preorder[1:len(leftinorder)]
			rightpreorder = preorder[1+len(leftpreorder):]
			break
		}
	}
	tree.Left = BuildTree(leftpreorder, leftinorder)
	tree.Right = BuildTree(rightpreorder, rightinorder)
	return tree
}
