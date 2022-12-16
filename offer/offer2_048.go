package offer

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor48() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var dfs func(*TreeNode)
	s := &strings.Builder{}
	dfs = func(tn *TreeNode) {
		if tn == nil {
			s.WriteString("nil,")
			return
		}
		s.WriteString(strconv.Itoa(tn.Val))
		s.WriteString(",")
		dfs(tn.Left)
		dfs(tn.Right)

	}
	dfs(root)
	return s.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	index := 0
	var build func() *TreeNode
	build = func() *TreeNode {
		if index >= len(sp) {
			return nil
		}
		if sp[index] == "nil" {
			index++
			return nil
		}
		val, _ := strconv.Atoi(sp[index])
		index++
		return &TreeNode{Val: val, Left: build(), Right: build()}
	}
	return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
