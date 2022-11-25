package offer

// 普通做法 读出listnode每个值然后反序
func ReversePrint(head *ListNode) []int {
	ans := make([]int, 0)
	ptr := head
	for {
		if ptr != nil {
			ans = append(ans, ptr.Val)
			ptr = ptr.Next
		} else {
			break
		}
	}
	ans2 := make([]int, 0)
	for i := len(ans) - 1; i >= 0; i-- {
		ans2 = append(ans2, ans[i])
	}
	return ans2
}

var ans []int

// 从后往前读可以利用栈的思想，压到最后回溯
func ReversePrint1(head *ListNode) []int {
	ans1 := make([]int, 0)
	Recur2(&ans1, head)
	return ans
}
func ReversePrint3(head *ListNode) []int {
	return Recur3(nil, head)
}
func Recur3(ans []int, head *ListNode) []int {
	if head == nil {
		return make([]int, 0)
	}
	ans1 := Recur3(ans, head.Next)
	ans1 = append(ans1, head.Val)
	return ans1

}

func Recur2(ans1 *[]int, head *ListNode) {
	if head == nil {
		return
	}
	Recur2(ans1, head.Next)
	*ans1 = append(*ans1, head.Val)
}

func Recur(head *ListNode) {
	if head == nil {
		return
	}
	Recur(head.Next)
	ans = append(ans, head.Val)
}
