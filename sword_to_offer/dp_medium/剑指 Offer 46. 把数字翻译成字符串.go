package dp_medium

import "strconv"

func translateNum(num int) int {
	s := strconv.Itoa(num)
	n := len(s)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		sub := s[i-1 : i+1]
		if sub >= "10" && sub <= "25" {
			if i == 1 {
				dp[i] = dp[i-1] * 2
			} else {
				dp[i] = dp[i-1] + dp[i-2]
			}
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[n-1]
}
