package offer

func lenLongestFibSubseq(arr []int) int {
	var max func(int, int) int
	max = func(i1, i2 int) int {
		if i1 > i2 {
			return i1
		}
		return i2
	}
	hashmap := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		hashmap[arr[i]] = i
	}
	dp := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = make([]int, i+1)
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		for j := i; j >= 0 && arr[j]*2 > arr[i]; j-- {
			if k, ok := hashmap[arr[i]-arr[j]]; ok {
				dp[i][j] = max(dp[j][k]+1, 3)
				res = max(dp[i][j], res)
			}
		}
	}
	return res
}
