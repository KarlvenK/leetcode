package list_and_stack

import "container/list"

type MaxQueue struct {
	que   *list.List
	deque *list.List
}

/*
func Constructor() MaxQueue {
	return MaxQueue{
		list.New(),
		list.New(),
	}
}
*/

func (m *MaxQueue) Max_value() int {
	if m.deque.Len() == 0 {
		return -1
	}
	return m.deque.Front().Value.(int)
}

func (m *MaxQueue) Push_back(value int) {
	for m.deque.Len() != 0 && m.deque.Back().Value.(int) < value {
		m.deque.Remove(m.deque.Back())
	}
	m.que.PushBack(value)
	m.deque.PushBack(value)
}

func (m *MaxQueue) Pop_front() int {
	if m.que.Len() == 0 {
		return -1
	}
	ans := m.que.Front().Value.(int)
	if m.deque.Front().Value.(int) == ans {
		m.deque.Remove(m.deque.Front())
	}
	m.que.Remove(m.que.Front())
	return ans
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
