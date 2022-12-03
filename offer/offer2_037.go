package offer

func asteroidCollision(asteroids []int) []int {
	res := []int{}
	for _, v := range asteroids {
		for v < 0 && len(res) > 0 && res[len(res)-1] > 0 {
			sum := res[len(res)-1] + v
			if sum >= 0 {
				v = 0
			}
			if sum <= 0 {
				res = res[:len(res)-1]
			}
		}

		if v != 0 {
			res = append(res, v)
		}
	}
	return res
}
