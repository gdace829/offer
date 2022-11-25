package offer

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	head1 := new(Node)
	now1 := head1
	now := head
	nodemap := make(map[*Node]*Node)
	for head != nil {
		now1.Val = head.Val
		if head.Next != nil {
			now1.Next = new(Node)
		}
		nodemap[head] = now1
		head = head.Next
		now1 = now1.Next
	}
	now2 := head1
	for now != nil {
		now2.Random = nodemap[now.Random]
		now2 = now2.Next
		now = now.Next
	}
	return head1
}
