package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	var merge func(*ListNode, *ListNode) *ListNode
	merge = func(ln1, ln2 *ListNode) *ListNode {
		list := &ListNode{}
		cur := list
		for ln1 != nil && ln2 != nil {
			if ln1.Val > ln2.Val {
				cur.Next = ln2
				ln2 = ln2.Next
			} else {
				cur.Next = ln1
				ln1 = ln1.Next
			}
			cur = cur.Next
		}
		if ln1 == nil {
			cur.Next = ln2
		}
		if ln2 == nil {
			cur.Next = ln1
		}
		return list.Next
	}
	if head == nil || head.Next == nil {
		return head
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next, slow = nil, slow.Next
	return merge(sortList(head), sortList(slow))

}
