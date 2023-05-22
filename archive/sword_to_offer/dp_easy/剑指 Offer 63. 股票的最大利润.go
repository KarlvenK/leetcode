package dp_easy

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 0x7fffffff
	}
	dp[0] = prices[0]
	ans := 0
	for i := 1; i < n; i++ {
		dp[i] = min(dp[i-1], prices[i])
		ans = max(ans, prices[i]-dp[i])
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
