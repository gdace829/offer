package offer

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, v := range tokens {
		a, err := strconv.Atoi(v)
		if err == nil {
			stack = append(stack, a)
		} else {
			switch v {
			case "+":
				{
					res := stack[len(stack)-1] + stack[len(stack)-2]
					stack = stack[:len(stack)-2]
					stack = append(stack, res)
					break
				}
			case "-":
				{
					res := stack[len(stack)-2] - stack[len(stack)-1]
					stack = stack[:len(stack)-2]
					stack = append(stack, res)
					break
				}
			case "*":
				{
					res := stack[len(stack)-1] * stack[len(stack)-2]
					stack = stack[:len(stack)-2]
					stack = append(stack, res)
					break
				}
			case "/":
				{
					res := stack[len(stack)-2] / stack[len(stack)-1]
					fmt.Println(res)
					stack = stack[:len(stack)-2]
					stack = append(stack, res)
					break
				}
			}
		}

	}
	return stack[0]
}
