package main

import (
	"fmt"
	"math"
)

func main() {
	var nums []int
	nums = []int{1, 2, 3, 4}
	fmt.Println(find132pattern(nums))
	nums = []int{1, 0, 1, -4, -3}
	fmt.Println(find132pattern(nums))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func find132pattern(nums []int) bool {
	n := len(nums)
	maxK := math.MinInt32
	candinateK := make([]int, 0)

	for i := n - 1; i >= 0; i-- {
		if nums[i] < maxK {
			return true
		}
		for len(candinateK) > 0 && candinateK[len(candinateK)-1] < nums[i] {
			maxK = max(maxK, candinateK[len(candinateK)-1])
			candinateK = candinateK[:len(candinateK)-1]
		}
		candinateK = append(candinateK, nums[i])
	}

	return false
}

/*
给定一个整数序列：a1, a2, ..., an，一个132模式的子序列ai, aj, ak
被定义为：当 i < j < k 时，ai < ak < aj。
设计一个算法，当给定有 n 个数字的序列时，验证这个序列中是否含有132模式的子序列。
*/
