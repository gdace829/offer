package offer

type MyCalendar struct {
	root *Node2
}

func Constructor58() MyCalendar {
	return MyCalendar{root: &Node2{}}
}

func (m *MyCalendar) Book(start int, end int) bool {
	if query(m.root, 0, 1e9, start, end) != 0 {

		return false
	} else {
		update(m.root, 0, 1e9, start, end, 1)
		return true
	}
}

type Node2 struct {
	left, right *Node2
	val         int
	add         int
}

func update(node *Node2, start int, end int, l int, r int, val int) {
	if l <= start && r >= end {
		node.val += val
		node.add += val
		return
	}
	if node.left == nil {
		node.left = &Node2{}
	}
	if node.right == nil {
		node.right = &Node2{}
	}
	if node.add != 0 {
		node.left.val += node.add
		node.right.val += node.add
		node.left.add += node.add
		node.right.add += node.add
		node.add = 0
	}
	mid := (start + end) >> 1
	if l <= mid {
		update(node.left, start, mid, l, r, val)
	}
	if r > mid {
		update(node.right, mid+1, end, l, r, val)
	}
	if node.left.val > node.right.val {
		node.val = node.left.val
	} else {
		node.val = node.right.val
	}

}
func query(node *Node2, start int, end int, l int, r int) int {
	if l <= start && r >= end {
		return node.val
	}
	if node.left == nil {
		node.left = &Node2{}
	}
	if node.right == nil {
		node.right = &Node2{}
	}
	if node.add != 0 {
		node.left.val += node.add
		node.right.val += node.add
		node.left.add += node.add
		node.right.add += node.add
		node.add = 0
	}
	mid := (start + end) >> 1
	ans := 0
	if l <= mid {
		ans = query(node.left, start, mid, l, r)
	}
	if r > mid {
		if a := query(node.right, mid+1, end, l, r); a > ans {
			ans = a
		}
	}
	return ans
}
