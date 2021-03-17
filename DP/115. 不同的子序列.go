package main

import "fmt"

func main() {
	s := "rabbbit"
	t := "rabbit"
	fmt.Println(numDistinct(s, t))
}

func numDistinct(s string, t string) int {
	n, m := len(s), len(t)
	// m = len(t)   n = len(s)
	if m > n {
		return 0
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}
