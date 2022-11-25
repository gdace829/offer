package offer

// 字符串操作可以转化为byte数组处理
func ReplaceSpace(s string) string {
	str := make([]byte, 0)
	s1 := []byte(s)
	for i := 0; i < len(s1); i++ {
		if s1[i] == ' ' {
			str = append(str, '%', '2', '0')
		} else {
			str = append(str, s1[i])
		}

	}
	return string(str)
}
