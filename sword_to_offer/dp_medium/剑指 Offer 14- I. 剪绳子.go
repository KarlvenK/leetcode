package dp_medium

func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 4; i <= n; i++ {
		tmp := 0
		for j := 1; j < i; j++ {
			tmp = max(tmp, dp[i-j]*dp[j])
		}
		dp[i] = tmp
	}
	return dp[n]
}
