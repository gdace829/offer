package offer

import "strconv"

func restoreIpAddresses(s string) []string {
	var isfuhe func(string) bool
	isfuhe = func(s string) bool {
		if s[0] == '0' && len(s) > 1 {
			return false
		}
		if i, _ := strconv.Atoi(s); i > 255 {
			return false
		}
		return true
	}
	res := []string{}
	cur := []string{}
	var dfs func(int)
	dfs = func(idx int) {
		if len(cur) > 4 || (len(cur) == 4 && idx < len(s)) || (len(cur) < 4 && idx == len(s)) {
			return
		}
		if idx == len(s) {
			str := ""
			for i := 0; i < len(cur)-1; i++ {
				str = str + cur[i] + "."
			}
			str += cur[len(cur)-1]
			res = append(res, str)
		}
		for i := idx + 1; i <= len(s); i++ {
			if isfuhe(s[idx:i]) {
				cur = append(cur, s[idx:i])
				dfs(i)
				cur = cur[:len(cur)-1]
			}
		}
	}
	dfs(0)
	return res

}
