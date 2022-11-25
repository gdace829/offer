package offer

// 两个栈维护
type MinStack struct {
	nums      []int
	mins      []int
	minscount int
	count     int
}

/** initialize your data structure here. */
func Constructor2() MinStack {
	stack := new(MinStack)
	stack.nums = make([]int, 10000)
	stack.count = 0
	stack.mins = make([]int, 10000)
	stack.minscount = 0
	return *stack
}

func (this *MinStack) Push(x int) {
	this.nums[this.count] = x
	this.count++
	if this.minscount == 0 {
		this.mins[this.minscount] = x
		this.minscount++
	} else if x <= this.mins[this.minscount-1] {
		this.mins[this.minscount] = x
		this.minscount++
	}
}

func (this *MinStack) Pop() {
	if this.count == 0 || this.minscount == 0 {
		return
	}
	if this.nums[this.count-1] == this.mins[this.minscount-1] {
		this.mins[this.minscount-1] = 0
		this.minscount--
	}
	this.nums[this.count-1] = 0
	this.count--
}

func (this *MinStack) Top() int {
	return this.nums[this.count-1]
}

func (this *MinStack) Min() int {
	if this.minscount == 0 {
		return 0
	}
	return this.mins[this.minscount-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
