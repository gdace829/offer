package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	stack []*TreeNode
	tn    *TreeNode
}

func Constructor55(root *TreeNode) BSTIterator {
	return BSTIterator{tn: root, stack: []*TreeNode{}}
}

func (this *BSTIterator) Next() int {
	for this.tn != nil {
		this.stack = append(this.stack, this.tn)
		this.tn = this.tn.Left
	}

	this.tn = this.stack[len(this.stack)-1]
	this.stack = this.stack[0 : len(this.stack)-1]
	res := this.tn.Val
	if this.tn.Right != nil {
		this.tn = this.tn.Right
	} else {
		this.tn = nil
	}
	return res
}

func (this *BSTIterator) HasNext() bool {
	return this.tn != nil || len(this.stack) != 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
