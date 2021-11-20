package string

import "strings"

func strToInt(s string) int {
	a, b := (1<<30)-1+(1<<30), -(1 << 31)
	s = strings.TrimSpace(s)
	if len(s) == 0 || (s[0] != '-' && s[0] != '+' && (s[0] < '0' || s[0] > '9')) {
		return 0
	}
	ans := 0
	sig := 1
	start := 0
	if s[0] == '-' || s[0] == '+' {
		start = 1
		if s[0] == '-' {
			sig = -1
		}
	}
	for i := start; i < len(s); i++ {
		if ans > a {
			break
		}
		if s[i] < '0' || s[i] > '9' {
			break
		} else {
			ans = ans*10 + int(s[i]-'0')
		}
	}
	ans *= sig
	if ans > a {
		ans = a
	} else if ans < b {
		ans = b
	}
	return ans
}
