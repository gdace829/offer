package offer

type CQueue struct {
	stack  *Stack
	stack1 *Stack
}

func Constructor1() CQueue {
	cqueue := new(CQueue)
	cqueue.stack = new(Stack)
	cqueue.stack1 = new(Stack)
	cqueue.stack.nums = make([]int, 1000)
	cqueue.stack1.nums = make([]int, 1000)
	return *cqueue
}

func (this *CQueue) AppendTail(value int) {
	this.stack.push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack.len == 0 && this.stack1.len == 0 {
		return -1
	}
	if this.stack1.len != 0 {
		return this.stack1.pop()
	} else {
		for {
			if this.stack.len != 0 {
				t := this.stack.pop()
				this.stack1.push(t)
			} else {
				break
			}

		}
		return this.stack1.pop()
	}

}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
