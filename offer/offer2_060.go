package offer

import "container/heap"

type minHeap1 struct {
	nums [][]int
}

func (min minHeap1) Len() int {
	return len(min.nums)
}
func (min minHeap1) Less(i, j int) bool {
	return min.nums[i][1] > min.nums[j][1]
}
func (min minHeap1) Swap(i, j int) {
	min.nums[i], min.nums[j] = min.nums[j], min.nums[i]
}
func (min *minHeap1) Push(x interface{}) {
	(*min).nums = append((*min).nums, x.([]int))
}
func (min *minHeap1) Pop() interface{} {
	old := (*min).nums
	n := len(old)
	x := old[n-1]
	(*min).nums = old[0 : n-1]
	return x
}
func topKFrequent(nums []int, k int) []int {
	m := &minHeap1{nums: [][]int{}}
	hashmap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		hashmap[nums[i]]++
	}
	for k, v := range hashmap {
		heap.Push(m, []int{k, v})
	}
	i := 0
	res := []int{}
	for i < k {
		res = append(res, heap.Pop(m).([]int)[0])
		i++
	}
	return res
}
