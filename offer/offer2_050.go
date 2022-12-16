package offer

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	treemap := map[int]int{}
	cur := 0
	res := 0
	var dfs func(*TreeNode)
	dfs = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		cur += tn.Val
		treemap[cur]++
		res += treemap[cur-targetSum]
		dfs(tn.Left)
		dfs(tn.Right)
		treemap[cur]--
		cur -= tn.Val
	}
	dfs(root)
	return res

}
