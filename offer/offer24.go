package offer

func reverseList2(head *ListNode) *ListNode {
	pre := head
	now := head
	next := head
	if now.Next == nil {
		return now
	}
	now = now.Next
	next = now.Next
	now.Next = pre
	pre.Next = nil
	pre = now
	now = next
	if now == nil {
		return pre
	}
	next = now.Next
	for {
		if now != nil {
			now.Next = pre
			pre = now
			now = next
			if now == nil {
				return pre
			} else {
				next = now.Next
			}
		}
	}

}

// 递归求解更方便
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}
