package offer

func majorityElement(nums []int) int {
	count := len(nums)
	numsmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsmap[nums[i]]++
		if numsmap[nums[i]] > count/2 {
			return numsmap[nums[i]]
		}
	}
	return -1
}
