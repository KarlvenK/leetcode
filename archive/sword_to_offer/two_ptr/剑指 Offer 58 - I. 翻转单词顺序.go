package two_ptr

import "strings"

func reverseWords(s string) string {
	list := strings.Split(s, " ")
	res := make([]string, 0)
	for i := len(list) - 1; i >= 0; i-- {
		s := strings.TrimSpace(list[i])
		if len(s) > 0 {
			res = append(res, s)
		}
	}
	return strings.Join(res, " ")
}
