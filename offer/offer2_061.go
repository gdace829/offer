package offer

import "container/heap"

type maxheap61 [][]int

func (max *maxheap61) Push(x interface{}) {
	*max = append(*max, x.([]int))
}
func (max *maxheap61) Pop() interface{} {
	old := *max
	length := len(old)
	x := old[length-1]
	*max = old[:length-1]
	return x
}

func (max maxheap61) Len() int           { return len(max) }
func (max maxheap61) Swap(i, j int)      { max[i], max[j] = max[j], max[i] }
func (max maxheap61) Less(i, j int) bool { return max[i][0]+max[i][1] > max[j][0]+max[j][1] }
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	maxh := &maxheap61{}
	for i := 0; i < len(nums1) && i < k; i++ {
		for j := 0; j < len(nums2) && j < k; j++ {
			heap.Push(maxh, []int{nums1[i], nums2[j]})
			if maxh.Len() > k {
				heap.Pop(maxh)
			}
		}

	}
	return *maxh

}
