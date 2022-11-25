package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	headA1, headB1 := headA, headB
	for headA1 != headB1 { //利用循环后长度相同点处理
		if headA1 == nil {
			headA1 = headB

		} else {
			headA1 = headA1.Next
		}
		if headB1 == nil {
			headB1 = headA

		} else {
			headB1 = headB1.Next
		}

	}
	return headA1
}
