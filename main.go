package main

import (
	"fmt"

	"sjs.com/m/competition"
)

func main() {
	// nums := make([]int, 5)
	// nums = append(nums, 1)
	// fmt.Println(nums[5])
	// fmt.Println("sjs")
	strs := []string{"aabb", "ab", "ba"}
	fmt.Println(competition.SimilarPairs(strs))

}
