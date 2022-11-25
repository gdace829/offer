package offer

func subarraySum(nums []int, k int) int {
	numsmap := make(map[int]int)
	numsmap[0]++
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum-k >= 0 {
			res += numsmap[sum-k]
		}
		numsmap[sum]++

	}
	return res
}
