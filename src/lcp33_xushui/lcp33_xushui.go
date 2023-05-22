package lcp33_xushui

import "container/heap"

type pair struct {
	tot int
	idx int
}

type Heap []pair

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].tot > h[j].tot
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func calc(a, b int) int {
	return (a + b - 1) / b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func storeWater(bucket []int, vat []int) int {
	n := len(bucket)
	h := make(Heap, 0)
	hp := &h
	heap.Init(hp)
	cnt := 0
	for i := 0; i < n; i++ {
		if vat[i] == 0 {
			continue
		}
		if bucket[i] == 0 {
			cnt++
			bucket[i]++
		}
		heap.Push(hp, pair{
			tot: calc(vat[i], bucket[i]),
			idx: i,
		})
	}

	if hp.Len() == 0 {
		return 0
	}

	ans := 1234567890

	for hp.Len() > 0 {
		t := h[0]
		heap.Pop(hp)
		if cnt > ans {
			break
		}
		ans = min(ans, t.tot+cnt)
		bucket[t.idx]++
		heap.Push(hp, pair{
			tot: calc(vat[t.idx], bucket[t.idx]),
			idx: t.idx,
		})
		cnt++
	}
	if ans == 1234567890 {
		return 0
	}
	return ans
}
