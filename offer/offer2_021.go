package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	headfirst:=&ListNode{}
	headfirst.Next=head
	deleteFromEnd(headfirst, n)
	return headfirst.Next
}

func deleteFromEnd(cur *ListNode, n int) int {
	if cur == nil {
		return 0
	}
	i := deleteFromEnd(cur.Next, n)
	if i == n {
		cur.Next = cur.Next.Next
	}
	return i + 1

}
