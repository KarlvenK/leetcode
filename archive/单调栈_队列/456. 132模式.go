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
给定一个整数序列：a1, a2, ..., an，一个132模式的子序列ai, aj, ak 满足：i < j < k， 且 ai < ak < aj。
设计一个算法，当给定有 n 个数字的序列时，验证这个序列中是否含有132模式的子序列。

solution:
	从右向左枚举 i, j, k 中的i， 我们枚举到i的时候必然已经遍历过j和k了
	所以问题就变成如何在之前已经遍历过的数中找到aj和ak
	使用单调栈，栈顶为最小元素， 同时我们要维护一个 maxK表示所有满足的k中最大的ak
	只要ai比ak小那么就找到了那么一组i, j, k, ai < ak < aj

	对单调栈的解释：
		对于当前的ai（还未入栈），在循环中的第一个if没有通过，说明它无法成为ai，那么只能成为aj或者ak
		栈内元素下标必然大于i，那么如果ai大于栈顶元素，那么栈顶元素就是一个符合的ak，我们只要记录符合要求的最大的ak
		所以用到了maxK变量
*/
