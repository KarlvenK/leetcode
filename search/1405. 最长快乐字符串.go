package search

import "sort"

func longestDiverseString(na int, nb int, nc int) string {
	ans := make([]byte, 0)
	type node struct {
		num  int
		char byte
	}
	sets := []node{
		{na, 'a'},
		{nb, 'b'},
		{nc, 'c'},
	}

	for {
		sort.Slice(sets, func(i, j int) bool {
			return sets[i].num > sets[j].num
		})
		mark := false
		for i, ch := range sets {
			if ch.num <= 0 {
				break
			}
			if len(ans) >= 2 && ch.char == ans[len(ans)-2] && ch.char == ans[len(ans)-1] {
				continue
			}
			mark = true
			ans = append(ans, ch.char)
			sets[i].num--
			break
		}

		if mark == false {
			break
		}
	}
	return string(ans)
}
