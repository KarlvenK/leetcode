package main

import "fmt"

func triangularSum(nums []int) int {
	temp := make([]int, len(nums))
	for t := 1; t <= len(nums)-1; t++ {
		for i := 0; i < len(nums)-1; i++ {
			temp[i] = nums[i] + nums[i+1]
			temp[i] %= 10
		}
		nums = append(make([]int, 0), temp...)
	}
	return nums[0]
}

func main() {
	nums := make([]int, 0)
	for i := 0; i < 1005; i++ {
		nums = append(nums, i)
	}
	fmt.Println(triangularSum(nums))
}
