package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	quick := head
	slow := head
	for quick != nil && quick.Next != nil {
		quick, slow = quick.Next.Next, slow.Next
		if quick == slow { //慢指针走的步数为nb
			for head != slow {
				head, slow = head.Next, slow.Next
			}
			return head
		}
	}
	return nil

}
