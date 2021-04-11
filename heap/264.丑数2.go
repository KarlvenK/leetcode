package binaryHeap

import "container/heap"

type node struct {
	data []int
}

func (x node) Len() int {
	return len(x.data)
}

func (x node) Less(i, j int) bool {
	return x.data[i] < x.data[j]
}

func (x node) Swap(i, j int) {
	x.data[i], x.data[j] = x.data[j], x.data[i]
}

func (x *node) Push(v interface{}) {
	x.data = append(x.data, v.(int))
}

func (x *node) Pop() interface{} {
	old := x.data
	v := old[len(old)-1]
	x.data = old[:len(old)-1]
	return v
}

func nthUglyNumber(n int) int {
	h := &node{[]int{1}}
	heap.Init(h)
	tool := make(map[int]int)
	tool[1] = 1

	var facs = []int{2, 3, 5}

	for i := 1; ; i++ {
		v := heap.Pop(h).(int)
		if i == n {
			return v
		}
		for _, fac := range facs {
			next := v * fac
			if tool[next] != 1 {
				tool[next] = 1
				heap.Push(h, next)
			}
		}
	}
}
