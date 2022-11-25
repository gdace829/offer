package offer

func pivotIndex(nums []int) int {
	rsum := 0
	for _, v := range nums {
		rsum += v
	}
	lsum := 0
	for index, v := range nums {
		rsum -= v
		if rsum == lsum {
			return index
		}
		lsum += v
	}
	return -1
}
