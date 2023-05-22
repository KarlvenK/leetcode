package main

import "fmt"

func partition(s string) [][]string {
	n := len(s) // the length of s
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, n)
	}
	var isPalindrome func(i, j int) int
	isPalindrome = func(i, j int) int {
		if f[i][j] != 0 {
			return f[i][j]
		}
		if i >= j {
			f[i][j] = 1
			return 1
		}
		if s[i] == s[j] {
			f[i][j] = isPalindrome(i+1, j-1)
		} else {
			f[i][j] = -1
		}
		return f[i][j]
	}
	/*

					true    			i >= j
		f[i][j] =
					f[i + 1][j - 1] 	i < j
	*/
	ans := make([][]string, 0)
	ss := make([]string, 0)
	var dfs func(index int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), ss...))
		}
		for j := i; j < n; j++ {
			if isPalindrome(i, j) == 1 {
				ss = append(ss, s[i:j+1])
				dfs(j + 1)
				ss = ss[:len(ss)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
