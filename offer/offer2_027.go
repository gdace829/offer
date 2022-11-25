package offer

func isPalindrome2(head *ListNode) bool {
	var getmid func(head *ListNode) *ListNode
	getmid = func(head *ListNode) *ListNode {
		quick := head
		slow := head
		for quick.Next != nil && quick.Next.Next != nil {
			slow = slow.Next
			quick = quick.Next.Next
		}
		return slow
	}
	var head2 *ListNode
	var relist func(head *ListNode) *ListNode
	relist = func(head *ListNode) *ListNode {
		if head.Next == nil {
			head2 = head
			return head
		}
		relist(head.Next).Next = head
		return head
	}
	mid := getmid(head)
	mid2 := mid.Next
	mid.Next = nil
	head2 = reverseList3(mid2)
	for head2 != nil && head != nil {
		if head.Val != head2.Val {
			return false
		}
		head2 = head2.Next
		head = head.Next
	}
	return true
}

func reverseList3(head *ListNode) *ListNode {

	var last *ListNode

	for head != nil {

		Nextnode := head.Next
		head.Next = last

		last = head
		head = Nextnode
	}
	return last
}
