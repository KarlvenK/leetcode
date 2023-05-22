package sorting

import "container/heap"

type queType []int

func (q queType) Len() int {
	return len(q)
}
func (q queType) Less(i, j int) bool {
	return q[i] > q[j]
}
func (q queType) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q *queType) Push(x interface{}) {
	*q = append(*q, x.(int))
}
func (q *queType) Pop() interface{} {
	old := *q
	x := old[len(old)-1]
	*q = old[:len(old)-1]
	return x
}
func getLeastNumbers(arr []int, k int) []int {
	que := &queType{}
	heap.Init(que)
	if k == 0 {
		return []int{}
	}
	for i := 0; i < k; i++ {
		heap.Push(que, arr[i])
	}
	for i := k; i < len(arr); i++ {
		if arr[i] < (*que)[0] {
			heap.Pop(que)
			heap.Push(que, arr[i])
		}
	}
	var ans []int
	for (*que).Len() != 0 {
		ans = append(ans, heap.Pop(que).(int))
	}
	return ans
}
