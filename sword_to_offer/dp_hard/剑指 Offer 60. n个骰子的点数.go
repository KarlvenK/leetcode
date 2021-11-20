package dp_hard

func dicesProbability(n int) []float64 {
	var dp [12][70]int

	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := i; j <= (6 * i); j++ {
			for k := 1; k <= 6; k++ {
				if j <= k {
					break
				}
				dp[i][j] += dp[i-1][j-k]
			}
		}
	}

	tot := pow(6, n)
	var ans []float64
	for i := n; i <= n*6; i++ {
		ans = append(ans, float64(dp[n][i])/float64(tot))
	}
	return ans
}
func pow(a, b int) int {
	ret := 1
	for i := 0; i < b; i++ {
		ret *= a
	}
	return ret
}
