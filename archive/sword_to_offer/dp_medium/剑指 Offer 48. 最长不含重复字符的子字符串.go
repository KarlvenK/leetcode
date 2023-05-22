package dp_medium

func lengthOfLongestSubstring(s string) int {
	str := []byte(s)
	marker := make(map[byte]struct{})
	i, j := 0, 0
	ans := 0
	n := len(str)
	for i <= j && j < n {
		if _, ok := marker[str[j]]; !ok {
			ans = max(ans, j-i+1)
			marker[str[j]] = struct{}{}
			j++
		} else {
			delete(marker, str[i])
			i++
		}
	}

	return ans
}
