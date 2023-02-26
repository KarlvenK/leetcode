package greedy

import (
	"container/heap"
	"sort"
)

type comb struct {
	id  int
	pft int
}
type comHeap []comb

func (h comHeap) Len() int           { return len(h) }
func (h comHeap) Less(i, j int) bool { return h[i].pft > h[j].pft }
func (h comHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *comHeap) Push(x interface{}) {
	*h = append(*h, x.(comb))
}
func (h *comHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	token := make([]int, n, n)
	for i := 0; i < n; i++ {
		token[i] = i
	}
	sort.Slice(token, func(i, j int) bool {
		return capital[token[i]] < capital[token[j]]
	})
	cur := 0

	h := &comHeap{}
	heap.Init(h)

	for cnt := 0; cnt < k; cnt++ {
		for cur < n && capital[token[cur]] <= w {
			heap.Push(h, comb{
				id:  token[cur],
				pft: profits[token[cur]],
			})
			cur++
		}
		if h.Len() == 0 {
			break
		}
		w += heap.Pop(h).(comb).pft
	}
	return w
}
