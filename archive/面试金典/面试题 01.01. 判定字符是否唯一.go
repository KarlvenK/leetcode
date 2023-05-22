package interview

func isUnique(astr string) bool {
	dict := make(map[rune]struct{})
	for _, ch := range astr {
		if _, ok := dict[ch]; ok {
			return false
		}
		dict[ch] = struct{}{}
	}
	return true
}
