package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	pre:=new(ListNode)
	now := head
	pre.Next = now
	for {
		if now == nil {
			break
		}
		if now.Val == val {
			pre.Next = now.Next
			break
		}
		now = now.Next
		pre = pre.Next
	}
	if head.Val==val {
		return pre.Next
	}else{
		return head
	}
}
