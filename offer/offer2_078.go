package offer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var merge func(*ListNode, *ListNode) *ListNode
	merge = func(ln1, ln2 *ListNode) *ListNode {
		head := &ListNode{}
		cur := head
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
		return head.Next
	}
	length := len(lists)
	if length == 0 {
		return nil
	}
	for i := 1; i < length; i *= 2 {
		for j := 0; j < length; j += i * 2 {
			if j+i < length {
				lists[j] = merge(lists[j], lists[j+i])
			}

		}
	}
	return lists[0]
}
