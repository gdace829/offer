package offer

func getKthFromEnd(head *ListNode, k int) *ListNode {
	pre := head
	now := head
	for {
		if k == 0 {
			pre = pre.Next
			now = now.Next
		} else {
			now = now.Next
			k--
		}
		if now == nil {
			return pre
		}
	}
}
