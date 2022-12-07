package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	root   *TreeNode
	leaves []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	cbt := CBTInserter{}
	cbt.root = root
	cbt.leaves = []*TreeNode{}
	cbt.leaves = append(cbt.leaves, root)
	for len(cbt.leaves) != 0 {
		if cbt.leaves[0].Left == nil {
			break
		}
		cbt.leaves = append(cbt.leaves, cbt.leaves[0].Left)
		if cbt.leaves[0].Right == nil {
			break
		}
		cbt.leaves = append(cbt.leaves, cbt.leaves[0].Right)
		cbt.leaves = cbt.leaves[1:]
	}
	return cbt
}

func (this *CBTInserter) Insert(v int) int {
	var res int
	node := &TreeNode{Val: v}
	if len(this.leaves) > 0 {
		res = this.leaves[0].Val
		if this.leaves[0].Left == nil {
			this.leaves[0].Left = node

		} else {
			this.leaves[0].Right = node
			this.leaves = this.leaves[1:]
		}
	}
	this.leaves = append(this.leaves, node)
	return res
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
