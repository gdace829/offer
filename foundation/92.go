package foundation

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	pre := &ListNode{}
	cur := &ListNode{}
	pre.Next = cur
	cur.Next = head
	head1 := cur
	for i := 0; i < left; i++ {
		pre = pre.Next
		cur = cur.Next
	}
	start := pre
	last := cur
	for j := left; j <= right; j++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	start.Next = pre
	last.Next = cur
	return head1.Next
}
