package offer

type MovingAverage struct {
	queue []int
	sum   int
	size  int
}

/** Initialize your data structure here. */
func Constructor41(size int) MovingAverage {
	return MovingAverage{queue: make([]int, 0, size), size: size}
}

func (this *MovingAverage) Next(val int) float64 {
	this.queue = append(this.queue, val)
	this.sum += val
	if len(this.queue) > this.size {
		this.sum -= this.queue[0]
		this.queue = this.queue[1:]
	}
	return float64(this.sum) / float64(len(this.queue))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
