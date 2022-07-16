package main

type WordFilter struct {
	dict map[string] int
}


func Constructor(words []string) WordFilter {
	wf := WordFilter {
		dict: make(map[string]int),
	}

	for i, word := range words {
		for j, n := 1, len(word); j <= n; j++ {
			for k := 0; k < n; k++ {
				wf.dict[word[:j] + "-" + word[k:]] = i
			}
		}
	}

	return wf
}


func (this *WordFilter) F(pref string, suff string) int {
	if idx, ok := this.dict[pref + "-" + suff]; ok {
		return idx
	}
	return -1
}


/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */

func main() {


}