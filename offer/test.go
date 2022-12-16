package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	first := &ListNode{}
	first.Next = head
	pre := first
	cur := first.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			cur = cur.Next
			continue
		}
		cur = cur.Next
		pre = pre.Next
	}
	return first.Next
}
