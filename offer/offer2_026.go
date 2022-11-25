package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	l1 := head
	l2 := head
	var getmid func(l1 *ListNode, l2 *ListNode) *ListNode
	getmid = func(l1, l2 *ListNode) *ListNode {
		for l2.Next != nil && l2.Next.Next != nil {
			l1 = l1.Next
			l2 = l2.Next.Next
		}
		return l1
	}
	mid := getmid(l1, l2)
	var head1 *ListNode
	var relist func(l1 *ListNode) *ListNode
	relist = func(l1 *ListNode) *ListNode {
		if l1.Next == nil {
			head1 = l1
			return l1
		}
		relist(l1.Next).Next = l1
		return l1
	}
	var mergelist func(l1 *ListNode, l2 *ListNode)
	mergelist = func(l1, l2 *ListNode) {
		for l1 != nil && l2 != nil {
			tmp1 := l1.Next
			tmp2 := l2.Next
			l1.Next = l2
			l1 = tmp1
			l2.Next = l1
			l2 = tmp2

		}

	}
	l1 = head
	l2 = mid.Next
	mid.Next = nil
	if l2 != nil {
		relist(l2).Next = nil
	}

	mergelist(l1, head1)
}
