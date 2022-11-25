package offer

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		aNode = &Node{Val: x}
		aNode.Next = aNode
		return aNode
	}
	cur := aNode
	next := aNode.Next
	head := aNode
	for aNode.Val <= next.Val {
		aNode = aNode.Next
		next = next.Next
		if next == cur {
			break
		}
	} //找到真正头节点,可以避免遍历不到head前面结点
	cur = next
	for x > next.Val {
		next = next.Next
		aNode = aNode.Next
		if next == cur {
			break
		}
	}
	t := &Node{Val: x}
	t.Next = next
	aNode.Next = t
	return head
}
