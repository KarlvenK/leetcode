package main

import (
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	maxSize, maxVal := 1, 1
	for i := 1; i < n; i++ {
		for j, v := range nums[:i] {
			if nums[i]%v == 0 {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
		if maxSize < dp[i] {
			maxSize, maxVal = dp[i], nums[i]
		}
	}

	ans := make([]int, 0)

	if maxSize == 1 {
		return []int{nums[0]}
	}

	for i := n - 1; i >= 0 && maxSize > 0; i-- {
		if dp[i] == maxSize && maxVal%nums[i] == 0 {
			ans = append(ans, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return ans
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
