package dp_hard

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	check := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		f[i] = make([]bool, n+1)
	}

	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if check(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else {
				if check(i, j) {
					f[i][j] = f[i][j] || f[i-1][j-1]
				}
			}
		}
	}
	return f[m][n]
}

/*
用dp算法
f[i][j] 表示s的前i个字符是否和p的前j个字符匹配
f[0][0] 表示空字符串匹配

对于 f[i][j]
若 p[j - 1] = '*': （两种情况）
	1. *匹配 0 次：
			那么f[i][j] 继承状态f[i][j - 2]
			(因为*和它前面的字符绑定，匹配零次相当于这两个字符消失)
	2.*匹配 1 次：
			前提是s[i - 1]要和 * 前的符号相等
			那么f[i][j] 继承状态f[i - 1][j]

*/
