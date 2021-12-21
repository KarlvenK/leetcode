package binaryHeap

import (
	"container/heap"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeap []*ListNode

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	dummyHead := &ListNode{Val: -1, Next: nil}

	h := new(minHeap)
	heap.Init(h)
	for _, v := range lists {
		if v != nil {
			heap.Push(h, v)
		}
	}
	pre := dummyHead
	for h.Len() != 0 {
		cur := heap.Pop(h).(*ListNode)
		pre.Next = cur
		pre = cur
		if cur.Next != nil {
			heap.Push(h, cur.Next)
		}
	}
	return dummyHead.Next
}
