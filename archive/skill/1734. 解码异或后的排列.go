package main

func decode(encoded []int) []int {
	n := len(encoded)
	total := 0
	for i := 1; i <= n; i++ {
		total = total ^ i
	}
	odd := 0
	for i := 1; i < n; i += 2 {
		odd = odd ^ encoded[i]
	}
	ans := make([]int, n+1)
	ans[0] = total ^ odd
	for i := 1; i <= n; i++ {
		ans[i] = ans[i-1] ^ encoded[i-1]
	}
	return ans
}
