package main

import (
	"fmt"
)

func minCut(s string) int {
	n := len(s)
	f := make([][]int, n)
	for i := range f {
		f[i] = make([]int, n)
	}

	var isPalindrome func(i, j int) int
	isPalindrome = func(i, j int) int {
		if i >= j {
			return 1
		}
		if f[i][j] != 0 {
			return f[i][j]
		}
		if s[i] == s[j] {
			f[i][j] = isPalindrome(i+1, j-1)
		} else {
			f[i][j] = -1
		}
		return f[i][j]
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = 100000
	}

	for i := 0; i < n; i++ {
		if isPalindrome(0, i) == 1 {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if isPalindrome(j+1, i) == 1 {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

func main() {
	s := "aab"
	fmt.Println(minCut(s))
}
