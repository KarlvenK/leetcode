package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	stack := make([]int, 0)
	size := 0
	for i := 0; i < (n << 1); i++ {
		for size != 0 && nums[i%n] > nums[stack[size-1]] {
			if ans[stack[size-1]] == -1 {
				ans[stack[size-1]] = nums[i%n]
			}
			stack = stack[:size-1]
			size--
		}
		stack = append(stack, i%n)
		size++
	}
	return ans
}

func main() {
	test := []int{1, 2, 3, 4, 3}
	fmt.Println(nextGreaterElements(test))
}
