package foundation

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	cur := head
	prehead := &ListNode{}
	prehead.Next = head
	pre := prehead
	//计数
	for cur != nil {
		cur = cur.Next
		count++
	}
	cur = head
	for count/k > 0 {
		curcount := 0
		start := pre
		last := cur
		for curcount < k {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
			curcount++
			count--
		}
		start.Next = pre
		last.Next = cur
		pre = last
	}
	return prehead.Next

}
