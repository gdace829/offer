package hot100

//O(n^2)
func twoSum(nums []int, target int) []int {
  for i := 0; i < len(nums); i++ {
	for j := i+1; j < len(nums); j++ {
		if nums[i]+nums[j]==target {
			return []int{i,j}
		}
	}
  }
  return []int{}
}
//better O(n)
func twoSum1(nums []int, target int) []int {
	m:=make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if index,ok:=m[target-nums[i]];ok {
			return []int{index,i}
		}
		m[nums[i]]=i
	 
	}
	return []int{}
}