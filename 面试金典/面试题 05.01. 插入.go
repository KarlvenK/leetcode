package interview
func insertBits(N int, M int, i int, j int) int {
	pm := M << (i)
	pl :=N >> (j + 1)
	pl = pl << (j + 1)
	pr := N & (pow(i) - 1)
	return pl | pm | pr

}
func pow(n int) int {
	ans := 1
	for i := 0; i < n; i++ {
		ans *= 2
	}
	return ans
}
