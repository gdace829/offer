package offer

import (
	"strconv"
	"strings"
)

// 闲暇时候研究
// 状态机题目很复杂
func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	_, err := strconv.ParseFloat(s, 10)
	if err == nil {
		return true
	}
	isOk := strings.Contains(err.Error(), strconv.ErrRange.Error())
	return isOk
}
