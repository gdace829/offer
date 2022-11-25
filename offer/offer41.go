package offer

import (
	"container/heap"
	"sort"
)

type MedianFinder1 struct {
	nums []int
}

/** initialize your data structure here. */
func Constructor3() MedianFinder1 {
	return MedianFinder1{nums: make([]int, 0)}
}

func (this *MedianFinder1) AddNum3(num int) {
	this.nums = append(this.nums, num)
	sort.Ints(this.nums)
}

func (this *MedianFinder1) FindMedian3() float64 {
	if len(this.nums)%2 == 0 {
		return (float64(this.nums[len(this.nums)/2]) + float64(this.nums[len(this.nums)/2+1])) / 2

	} else {
		return float64(this.nums[len(this.nums)/2])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
type MedianFinder struct {
	maxheap MaxHeap
	minheap Minheap
}

/** initialize your data structure here. */
func Constructor4() MedianFinder {
	m := MedianFinder{maxheap: make([]int, 0), minheap: make([]int, 0)}
	heap.Init(&m.maxheap)
	heap.Init(&m.minheap)
	return m
}

func (this *MedianFinder) AddNum(num int) {
	if this.maxheap.Len() == this.minheap.Len() {
		heap.Push(&this.minheap, num)
		heap.Push(&this.maxheap, heap.Pop(&this.minheap))
	} else {
		heap.Push(&this.maxheap, num)
		heap.Push(&this.minheap, heap.Pop(&this.maxheap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxheap.Len() == this.minheap.Len() {
		return (float64(this.maxheap[0]) + float64(this.minheap[0])) / 2
	} else {
		return float64(this.maxheap[0])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
