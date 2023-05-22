package main

import "fmt"

func main() {
	s := "rabbbit"
	t := "rabbit"
	fmt.Println(numDistinct(s, t))
}

func numDistinct(s string, t string) int {
	sLen, tLen := len(s), len(t)
	// m = len(t)   n = len(s)
	if tLen > sLen {
		return 0
	}
	dp := make([][]int, tLen+1)
	for i := range dp {
		dp[i] = make([]int, sLen+1)
	}

	for i := 0; i < sLen+1; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < tLen+1; i++ {
		for j := 1; j < sLen+1; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[tLen][sLen]
}

/*
solution:
 dp[i][j] 表示 t 的前 i 个字符在 s 的前 j 个字符中的子序列个数
 所以dp数组要开 (sLen + 1) * (tLen + 1)
 当 t[i - 1] = s[j - 1] （即第t的第i个字符和s的第j个字符相等）有两种情况：
	1)： dp[i][j] 可以从 dp[i - 1][j - 1]继承过来（讲t[i - 1] s[j - 1]与它们前面的字符串连接）
 	2):	 也可以从dp[i][j - 1]继承过来(无视s[j - 1]， 仍然靠前s的前 j - 1 个字符来匹配 t的前i个字符)

 当t[i - 1] != s[j - 1]时，显然从dp[i][j - 1]继承

 边界条件：
	t的前 0 个字符可以被s任意匹配，且方案只有一种
	所以dp[0][:]全部赋值为1
*/
