package offer

import (
	"strconv"
	"strings"
)

func Serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	// write code here
	var dfs func(root *TreeNode)
	str := ""
	dfs = func(root *TreeNode) {
		if root == nil {
			str += "#,"
			return
		}
		str += strconv.Itoa(root.Val) + ","
		dfs(root.Left)
		dfs(root.Right)

	}
	dfs(root)
	return str
}
func Deserialize(s string) *TreeNode {
	if s == "" {
		return nil
	}
	// write code here
	var redfs func() *TreeNode
	strs := strings.Split(s, ",")
	redfs = func() *TreeNode {
		if strs[0] == "#" {
			strs = strs[1:]
			return nil
		}
		num, _ := strconv.Atoi(strs[0])
		strs = strs[1:]
		root := &TreeNode{Val: num}
		if len(strs) != 0 {
			root.Left = redfs()
		}
		if len(strs) != 0 {
			root.Right = redfs()
		}
		return root
	}
	return redfs()
}
