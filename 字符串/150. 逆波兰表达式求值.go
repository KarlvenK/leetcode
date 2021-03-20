package main

import (
	"fmt"
	"strconv"
)

func main() {
	ss := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(ss))
}
func evalRPN(tokens []string) int {
	nums := make([]int, 0)
	for _, s := range tokens {
		val, err := strconv.Atoi(s)
		if err == nil {
			nums = append(nums, val)
		} else {
			a := nums[len(nums)-1]
			b := nums[len(nums)-2]
			nums = nums[:len(nums)-2]
			switch s {
			case "+":
				nums = append(nums, b+a)
			case "-":
				nums = append(nums, b-a)
			case "*":
				nums = append(nums, b*a)
			default:
				nums = append(nums, b/a)
			}
		}
	}
	return nums[0]
}
