package list_and_stack

import "container/list"

type MinStack struct {
	s1 *list.List
	s2 *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		s1: list.New(),
		s2: list.New(),
	}
}

func (this *MinStack) Push(x int) {
	this.s1.PushBack(x)
	var ans int
	if this.s2.Len() != 0 {
		ans = this.s2.Back().Value.(int)
	} else {
		ans = 0x7fffffff
	}
	if x > ans {
		this.s2.PushBack(ans)
	} else {
		this.s2.PushBack(x)
	}
}

func (this *MinStack) Pop() {
	this.s1.Remove(this.s1.Back())
	this.s2.Remove(this.s2.Back())
}

func (this *MinStack) Top() int {
	return this.s1.Back().Value.(int)
}

func (this *MinStack) Min() int {
	return this.s2.Back().Value.(int)
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
