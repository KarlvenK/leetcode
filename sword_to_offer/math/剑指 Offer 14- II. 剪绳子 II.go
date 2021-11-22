package math

func cuttingRope(n int) int {
	if n < 4 {
		return n - 1
	}
	p := 1000000007
	ans := 1
	for n > 4 {
		ans = ans * 3 % p
		n -= 3
	}
	ans = ans * n % p
	return ans
}
