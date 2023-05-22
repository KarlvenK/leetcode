package main

func nthUglyNumber(n int) int {
	f := make([]int, n+1)
	f[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		n2, n3, n5 := 2*f[p2], 3*f[p3], 5*f[p5]
		f[i] = min(n2, min(n3, n5))
		if f[i] == n2 {
			p2++
		}
		if f[i] == n3 {
			p3++
		}
		if f[i] == n5 {
			p5++
		}
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
编写一个程序，找出第 n 个丑数。

丑数就是质因数只包含 2, 3, 5 的正整数
1是丑数
n <= 1690


solution:
	dp[1] = 1
	指针 p2, p3, p5表示下一个丑数是当前指针指向的丑数乘以对应的质因数
	dp[i] = min(dp[p2] * 2, dp[p3] * 3, dp[p5] * 5)
	可能出现乘完以后相等的情况，指针哦毒药加一（三个if都要判断，不适用else）
*/
