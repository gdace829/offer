package offer

func containsNearbyAlmostDuplicate(nums []int, k, t int) bool {
	var getId func(int, int) int
	getId = func(i1, i2 int) int {
		if i1 >= 0 {
			return i1 / i2
		}
		return (i1+1)/i2 - 1
	}
	var abs func(int) int

	abs = func(i1 int) int {
		if i1 > 0 {
			return i1
		} else {
			return 0 - i1
		}
	}
	hashmap := map[int]int{}
	for index, value := range nums {
		id := getId(value, t+1)
		if _, has := hashmap[id]; has {
			return true
		}
		if y, has := hashmap[id-1]; has && abs(value-y) <= t {
			return true
		}
		if y, has := hashmap[id+1]; has && abs(value-y) <= t {
			return true
		}
		hashmap[id] = value
		if index >= k {
			delete(hashmap, getId(nums[index-k], t+1))
		}
	}
	return false

}
