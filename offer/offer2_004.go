package offer

func singleNumber(nums []int) int {
	res := int32(0)
	for j := 0; j < 32; j++ {
		sum := 0
		for i := 0; i < len(nums); i++ {
			sum += (nums[i] >> j) & 1
		}
		res |= int32((sum % 3) << j)
	}
	return int(res)
}
