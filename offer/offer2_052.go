package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {
	var bst func(*TreeNode)
	isfirst := false
	var head *TreeNode
	var cur *TreeNode
	bst = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		bst(tn.Left)
		if tn.Left == nil && isfirst == false {
			head = &TreeNode{Val: tn.Val}
			cur = head
			isfirst = true

		} else {
			cur.Right = &TreeNode{Val: tn.Val}
			cur = cur.Right
		}

		bst(tn.Right)
	}
	bst(root)
	return head

}
