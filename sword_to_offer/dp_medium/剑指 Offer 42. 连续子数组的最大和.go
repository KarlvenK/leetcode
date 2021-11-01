package dp_medium

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = -123456789
	}
	dp[0] = nums[0]
	ans := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
