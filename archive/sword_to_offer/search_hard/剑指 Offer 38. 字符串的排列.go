package search_hard

func permutation(s string) (ret []string) {
	char := []byte(s)
	mark := make([]bool, len(char))
	n := len(char)

	temp := make([]byte, n)
	var ans []string
	var dfs func(int)
	dfs = func(step int) {
		if step == n {
			ans = append(ans, string(temp))
			return
		}
		for i := 0; i < n; i++ {
			if mark[i] == false {
				mark[i] = true
				temp[step] = char[i]
				dfs(step + 1)
				mark[i] = false
			}
		}
	}
	dfs(0)
	tool := make(map[string]struct{})
	for i := 0; i < len(ans); i++ {
		if _, ok := tool[ans[i]]; ok {
			continue
		} else {
			tool[ans[i]] = struct{}{}
			ret = append(ret, ans[i])
		}
	}
	return
}
