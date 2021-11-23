package interview

import "strings"

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s := s2 + s2
	return strings.Index(s, s1) != -1
}
