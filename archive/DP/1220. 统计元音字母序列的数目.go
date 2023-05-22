package main

func countVowelPermutation(n int) int {
	mod := 1000000007
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 5)
	}
	for i := 0; i < 5; i++ {
		dp[0][i] = 1
	}
	/*
	   a   0
	   e   1
	   i   2
	   o   3
	   u   4
	*/
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][1] + dp[i-1][4] + dp[i-1][2]
		dp[i][1] = dp[i-1][0] + dp[i-1][2]
		dp[i][2] = dp[i-1][1] + dp[i-1][3]
		dp[i][3] = dp[i-1][2]
		dp[i][4] = dp[i-1][2] + dp[i-1][3]
		for j := 0; j < 5; j++ {
			dp[i][j] %= mod
		}
	}
	ans := 0
	for i := 0; i < 5; i++ {
		ans += dp[n-1][i]
		ans %= mod
	}
	return ans
}
