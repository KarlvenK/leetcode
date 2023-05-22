package main

import "fmt"

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return getMax(nums[0], nums[1])
	}
	var try func(num []int) int
	try = func(num []int) int {
		dp1, dp2 := num[0], getMax(num[0], num[1])
		for _, v := range num[2:] {
			dp1, dp2 = dp2, getMax(dp1+v, dp2)
		}
		return dp2
	}
	ans := try(nums[:n-1])
	ans = getMax(ans, try(nums[1:]))
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	fmt.Println(rob(nums))
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
