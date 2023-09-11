package hot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	var sum, add int
	for l1 != nil || l2 != nil {
		var a, b int
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum = a + b + add
		cur.Val = sum % 10
		add = sum / 10
		if add != 0 || l1 != nil || l2 != nil {
			cur.Next = &ListNode{}
			cur = cur.Next
		}

	}

	if add != 0 {
		cur.Val = add
	}
	return head

}
