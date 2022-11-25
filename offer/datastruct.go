package offer

// 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 栈
type Stack struct {
	nums []int
	len  int
}

func (s *Stack) pop() int {
	t := s.nums[s.len-1]
	s.nums[s.len-1] = 0
	s.len--
	return t
}
func (s *Stack) push(num int) {
	s.nums[s.len] = num
	s.len++
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
type Node1 struct {
	Val   int
	Prev  *Node1
	Next  *Node1
	Child *Node1
}
type MaxHeap []int // 定义一个类型

func (h MaxHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h MaxHeap) Less(i, j int) bool { // 绑定less方法
	return h[i] > h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h MaxHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

type Minheap []int // 定义一个类型

func (h Minheap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h Minheap) Less(i, j int) bool { // 绑定less方法
	return h[i] < h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h Minheap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *Minheap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *Minheap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}
