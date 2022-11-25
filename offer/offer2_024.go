package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var headfirst *ListNode
	var relist func(head *ListNode) *ListNode
	relist = func(head *ListNode) *ListNode {
		if head.Next == nil {
			headfirst = head
			return head
		}
		relist(head.Next).Next = head
		return head
	}
	relist(head).Next = nil
	return headfirst
}
