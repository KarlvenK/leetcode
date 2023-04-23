package main

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n+1, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = 1000000000
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		sum, m := 0, 0
		for j := i - 1; j >= 0; j-- {
			sum += books[j][0]
			m = max(m, books[j][1])
			if sum > shelfWidth {
				break
			}
			dp[i] = min(dp[i], dp[j]+m)
		}
	}
	return dp[n]
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
