package foundation

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func findmid(head *ListNode) *ListNode {
	slow := head
	quick := head
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	return slow
}
func reorderList(head *ListNode) {
	mid := findmid(head)
	var pre *ListNode
	for mid != nil {
		next := mid.Next
		mid.Next = pre
		pre = mid
		mid = next
	}
	for pre.Next != nil {
		next := head.Next
		next1 := pre.Next
		head.Next = pre
		pre.Next = next
		head = next
		pre = next1
	}
}
