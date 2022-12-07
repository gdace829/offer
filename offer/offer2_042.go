package offer

type RecentCounter struct {
	nums  []int
	index int
}

func Constructor42() RecentCounter {
	return RecentCounter{nums: []int{}, index: 0}
}

func (this *RecentCounter) Ping(t int) int {
	this.index = 0
	this.nums = append(this.nums, t)
	for t-3000 > this.nums[this.index] {
		this.index++
	}
	this.nums = this.nums[this.index:]
	return len(this.nums)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
