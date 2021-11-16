package sorting

import "container/heap"

type maxHeap []int
type minHeap []int

func (m maxHeap) Len() int {
	return len(m)
}
func (m minHeap) Len() int {
	return len(m)
}
func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}
func (m minHeap) Less(i, j int) bool {
	return m[i] < m[j]
}
func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}
func (m *maxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}
func (m *minHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

type MedianFinder struct {
	minH *minHeap
	maxH *maxHeap
}

// Constructor initialize your data structure here.
func Constructor() MedianFinder {
	return MedianFinder{
		minH: new(minHeap),
		maxH: new(maxHeap),
	}
}

func (th *MedianFinder) AddNum(num int) {
	if th.minH.Len() == th.maxH.Len() {
		if th.minH.Len() == 0 || num > (*th.minH)[0] {
			heap.Push(th.minH, num)
		} else {
			heap.Push(th.maxH, num)
			t := heap.Pop(th.maxH).(int)
			heap.Push(th.minH, t)
		}
	} else {
		if num > (*th.minH)[0] {
			heap.Push(th.minH, num)
			t := heap.Pop(th.minH).(int)
			heap.Push(th.maxH, t)
		} else {
			heap.Push(th.maxH, num)
		}
	}
}

func (th *MedianFinder) FindMedian() float64 {
	if th.maxH.Len() == th.minH.Len() {
		return (float64((*th.minH)[0]) + float64((*th.maxH)[0])) / 2
	}
	return float64((*th.minH)[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
