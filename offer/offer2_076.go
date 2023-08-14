package offer

import "container/heap"

type minHeap76 struct {
	nums []int
}

func (m minHeap76) Len() int {
	return len(m.nums)
}
func (m minHeap76) Less(i, j int) bool {
	return m.nums[i] < m.nums[j]
}
func (m minHeap76) Swap(i, j int) {
	m.nums[i], m.nums[j] = m.nums[j], m.nums[i]
}
func (m *minHeap76) Pop() interface{} {
	m1 := m
	length := len(m1.nums)
	x := m1.nums[length-1]
	m1.nums = m1.nums[:length-1]
	return x
}
func (m *minHeap76) Push(x interface{}) {
	m.nums = append(m.nums, x.(int))
}
func findKthLargest(nums []int, k int) int {
	m := &minHeap76{nums: []int{}}
	for _, v := range nums {
		if m.Len() == k {
			heap.Push(m, v)
			heap.Pop(m)
		} else {
			heap.Push(m, v)
		}

	}
	return m.nums[0]
}
