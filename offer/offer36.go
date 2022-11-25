package offer

// 可以在函数里面用匿名函数，不用使用全局变量
func Convert(root *TreeNode) *TreeNode {
	var dfs func(cur *TreeNode)
	var pre, head *TreeNode
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		dfs(cur.Left)
		if pre != nil {
			pre.Right = cur
		} else {
			head = cur
		}
		cur.Left = pre
		pre = cur
		dfs(cur.Right)
	}
	dfs(root)
	return head
}

func Convertsjs(root *TreeNode) *TreeNode {
	var dfs func(cur *TreeNode)
	var pre, head *TreeNode
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		dfs(cur.Left)
		if pre == nil {
			head = cur
		} else {
			pre.Right = cur
		}
		cur.Left = pre
		pre = cur
		//:todo
		dfs(cur.Right)
	}
	pre.Right = head
	head.Left = pre
	dfs(root)
	return head

}
