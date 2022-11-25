package offer

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1 []int
	var stack2 []int
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	var head *ListNode
	count := 0
	for len(stack1) > 0 || len(stack2) > 0 || count > 0 {
		sum := count
		if len(stack1) > 0 {
			sum += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			sum += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		count = sum / 10
		node := &ListNode{Val: sum % 10}
		if head != nil {
			node.Next = head
		}
		head = node
	}
	return head

}
