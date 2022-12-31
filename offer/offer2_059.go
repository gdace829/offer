package offer

import "container/heap"

type minHeap []int

func (min minHeap) Len() int               { return len(min) }
func (min minHeap) Less(i int, j int) bool { return min[i] < min[j] }
func (min minHeap) Swap(i int, j int)      { min[i], min[j] = min[j], min[i] }
func (min *minHeap) Pop() interface{} {
	old := *min
	n := len(old)
	x := old[n-1]
	*min = old[0 : n-1]
	return x
}
func (min *minHeap) Push(x interface{}) {
	*min = append(*min, x.(int))
}

type KthLargest struct {
	min *minHeap
	k   int
}

func Constructor59(k int, nums []int) KthLargest {

	K := KthLargest{min: &minHeap{}, k: k}
	for i := 0; i < len(nums); i++ {
		K.Add(nums[i])
	}
	return K
}

func (this *KthLargest) Add(val int) int {
	if this.min.Len() < this.k || val > (*this.min)[0] {
		heap.Push(this.min, val)
	}
	if this.min.Len() > this.k {
		heap.Pop(this.min)
	}
	return (*this.min)[0]
}
