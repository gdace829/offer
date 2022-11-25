package offer

func flatten(root *Node1) *Node1 {
	cur := root
	for cur != nil {
		next := cur.Next
		for cur.Child != nil {
			child := flatten(cur.Child)
			cur.Child = nil
			tail := child
			for tail.Next != nil {
				tail = tail.Next
			}
			cur.Next = child
			child.Prev = cur
			tail.Next = next
			if next != nil {
				next.Prev = tail
			}
		}
		cur = next
	}
	return root
}
