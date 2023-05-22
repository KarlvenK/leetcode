package string

import (
	"strings"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	if strings.Count(s, "e")+strings.Count(s, "E") > 1 {
		return false
	}
	if strings.Count(s, ".") > 1 {
		return false
	}
	pos := max(max(-1, strings.Index(s, "e")), strings.Index(s, "E"))
	if pos != -1 {
		return (isDecimal(s[:pos]) || isInteger(s[:pos])) && isInteger(s[pos+1:])
	} else {
		return isDecimal(s) || isInteger(s)
	}
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func isInteger(s string) bool {
	if len(s) == 0 {
		return false
	}
	start := 0
	if s[0] == '-' || s[0] == '+' {
		if len(s) == 1 {
			return false
		}
		start = 1
	}
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
func isDecimal(s string) bool {
	if len(s) == 0 {
		return false
	}
	start := 0
	if s[0] == '+' || s[0] == '-' {
		if len(s) == 1 {
			return false
		}
		start = 1
	}
	if strings.Count(s, ".") != 1 {
		return false
	}
	pos := strings.Index(s, ".")
	s1, s2 := s[start:pos], s[pos+1:]
	check := func(t string) bool {
		for i := 0; i < len(t); i++ {
			if t[i] < '0' || t[i] > '9' {
				return false
			}
		}
		return true
	}
	if len(s1) == 0 && len(s2) == 0 {
		return false
	}
	if len(s1) == 0 {
		return check(s2)
	}
	if len(s2) == 0 {
		return check(s1)
	}
	return check(s1) && check(s2)
}
