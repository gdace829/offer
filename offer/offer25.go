package offer

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	now := head
	for {
		if l1 != nil && l2 != nil {
			if l1.Val > l2.Val {
				now.Next = l2
				l2 = l2.Next

			} else {
				now.Next = l1
				l1 = l1.Next
			}
			now = now.Next
		} else {
			break
		}
	}
	if l1 == nil {
		now.Next = l2
	} else {
		now.Next = l1
	}
	return head.Next
}
