package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i] > n || nums[i] <= 0 {
			nums[i] = n + 1
		}
	}
	for i := 0; i < n; i++ {
		if abs(nums[i]) <= n {
			if nums[abs(nums[i])-1] > 0 {
				nums[abs(nums[i])-1] *= -1
			}
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return n + 1
}

func main() {
	arr := []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(arr))
}
