package list_and_stack

import "container/heap"

type node struct {
	val int
	pos int
}
type maxHeap []node

func (m maxHeap) Len() int {
	return len(m)
}
func (m maxHeap) Less(i, j int) bool {
	return m[i].val > m[j].val
}
func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(node))
}
func (m *maxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) (ans []int) {
	if len(nums) == 0 {
		return
	}
	que := &maxHeap{}

	heap.Init(que)
	for i := 0; i < k; i++ {
		heap.Push(que, node{
			nums[i],
			i,
		})
	}

	left, right := 0, k-1
	for ; right < len(nums); func() {
		left++
		right++
		if right < len(nums) {
			heap.Push(que, node{
				nums[right],
				right,
			})
		}
	}() {
		for {
			tmp := (*que)[0]
			if tmp.pos >= left && tmp.pos <= right {
				ans = append(ans, tmp.val)
				break
			} else {
				heap.Pop(que)
			}
		}
	}
	return
}
