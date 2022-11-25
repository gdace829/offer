package offer

import (
	"fmt"
)

func translateNum(num int) int {
	str := fmt.Sprintf("%d", num)
	dp := make([]int, len(str))
	if len(str) == 1 {
		return 1
	} else if len(str) == 2 {
		s := str[0:2]
		if s <= "25" && s >= "1" {
			return 2
		} else {
			return 1
		}
	}
	s := str[0:2]
	dp[0] = 1
	if s <= "25" && s >= "1" {
		dp[1] = 2
	} else {
		dp[1] = 1
	}
	for i := 2; i < len(str); i++ {
		s := str[i-1 : i+1]
		if s <= "25" && s >= "1" {
			dp[i] = dp[i-2] + dp[i-1]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(str)-1]
}
