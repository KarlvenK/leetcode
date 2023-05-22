package dp_easy

func numWays(n int) int {
	mod := 1000000007
	a, b := 1, 2
	if n == 0 {
		return 1
	}
	if n < 3 {
		return n
	}
	for i := 3; i <= n; i++ {
		t := b
		b = a + b
		b = b % mod
		a = t
	}
	return b
}
