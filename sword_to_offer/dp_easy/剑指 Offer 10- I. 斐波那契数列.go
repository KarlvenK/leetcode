package dp_easy

func fib(n int) int {
	mod := 1000000007
	a, b := 0, 1
	if n < 2 {
		return n
	}
	for i := 2; i <= n; i++ {
		t := b
		b = a + b
		b = b % mod
		a = t
	}
	return b
}
