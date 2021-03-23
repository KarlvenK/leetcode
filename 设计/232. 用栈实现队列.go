package main

type MyQueue struct {
	instack, outstack *stack
}
type stack struct {
	size   int
	member []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		instack: &stack{
			size:   0,
			member: make([]int, 0),
		},
		outstack: &stack{
			size:   0,
			member: make([]int, 0),
		},
	}
}

func (obj *stack) push(val int) {
	obj.member = append(obj.member, val)
	obj.size++
}

func (obj *stack) empty() bool {
	if obj.size == 0 {
		return true
	}
	return false
}

func (obj *stack) pop() int {
	var res int
	if !obj.empty() {
		t := obj.size - 1
		obj.size--
		res = obj.member[t]
		obj.member = obj.member[:t]
	}
	return res
}
func (obj *stack) top() int {
	if !obj.empty() {
		return obj.member[obj.size-1]
	}
	return 0
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.instack.push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.outstack.empty() {
		for !this.instack.empty() {
			this.outstack.push(this.instack.pop())
		}
	}
	return this.outstack.pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.outstack.empty() {
		for !this.instack.empty() {
			this.outstack.push(this.instack.pop())
		}
	}
	return this.outstack.top()
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.instack.empty() && this.outstack.empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
