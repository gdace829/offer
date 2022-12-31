package competition

import "math"

func smallestValue(n int) int {
	var small func(int) int
	small = func(num int) int {
		max := (int)(math.Sqrt(float64(num)) + 1)
		res := 0
		for i := 2; i < max; i++ {

			for num%i == 0 {
				num /= i
				res += i
			}
		}
		if num > 1 {
			return res + num
		}
		return res
	}
	last := n
	n = small(n)
	for last != n {
		last = n
		n = small(n)
	}
	return n
}
