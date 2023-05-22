package seq_search

func firstUniqChar(s string) byte {
	dict := make(map[int32]int)
	for _, ch := range s {
		dict[ch-'a']++
	}
	for _, ch := range s {
		if dict[ch-'a'] == 1 {
			return byte(ch)
		}
	}
	return ' '
}
