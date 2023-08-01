package everyday

// 1247 交换字符使得字符串相同
func minimumSwap(s1 string, s2 string) int {
	length := len(s1)
	a := 0
	b := 0
	for i := 0; i < length; i++ {
		if s1[i] != s2[i] {
			if s1[i] == 'x' {
				a++
			} else {
				b++
			}
		}
	}
	if a%2 == 0 && b%2 == 0 {
		return a/2 + b/2
	} else if a%2 == 1 && b%2 == 1 {
		return (a-1)/2 + (b+1)/2 + 1
	} else {
		return -1
	}
}
