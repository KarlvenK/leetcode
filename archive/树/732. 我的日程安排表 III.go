package main

import "fmt"

type pair struct {
	num  int
	lazy int
}

type MyCalendarThree map[int]pair

func Constructor() MyCalendarThree {
	return make(map[int]pair)
}

func (cld MyCalendarThree) update(begin, end, l, r, idx int) {
	if r < begin || l > end {
		return
	}
	//fmt.Println(begin, end, l, r, idx)
	if begin <= l && r <= end {
		p := cld[idx]
		p.num++
		p.lazy++
		cld[idx] = p
	} else {
		mid := l + (r-l)/2
		cld.update(begin, end, l, mid, idx*2)
		cld.update(begin, end, mid+1, r, idx*2+1)
		p := cld[idx]
		p.num = max(cld[idx*2].num, cld[idx*2+1].num) + p.lazy
		cld[idx] = p
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (cld MyCalendarThree) Book(start int, end int) int {
	cld.update(start, end-1, 0, 1e9, 1)
	return cld[1].num
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */

func main() {
	obj := Constructor()
	opt := [][]int{
		{10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55},
	}
	for _, v := range opt {
		x := obj.Book(v[0], v[1])
		fmt.Println(x)
	}
}
