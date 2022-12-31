package offer

import (
	"math/rand"
	"sort"
	"time"
)

type Solution struct {
	nums    []int
	presums []int
}

func Constructor(w []int) Solution {
	n := len(w)
	presum := make([]int, n)
	if n >= 0 {
		presum[0] = w[0]
	}
	for i := 1; i < n; i++ {
		presum[i] = presum[i-1] + w[i]
	}
	return Solution{nums: w, presums: presum}
}

func (this *Solution) PickIndex() int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(this.presums[len(this.presums)-1])
	return sort.SearchInts(this.presums, num+1)
}
