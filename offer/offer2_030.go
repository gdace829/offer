package offer

import "math/rand"

type RandomizedSet struct {
	nums []int
	dict map[int]int
}

/** Initialize your data structure here. */
func Constructor6() RandomizedSet {
	return RandomizedSet{nums: make([]int, 0), dict: make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dict[val]; ok {
		return false
	}
	this.dict[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.dict[val]
	if !ok {
		return false
	}
	this.nums[id] = this.nums[len(this.nums)-1]
	this.dict[this.nums[id]] = id
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.dict, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
