package sword_to_offer

type CQueue struct {
	s1, s2 stack
}

type stack struct {
	sz   int
	data []int
}

func (s *stack) Push(x int) {
	s.sz++
	s.data = append(s.data, x)
}

func (s *stack) Pop() (ret int) {
	s.sz--
	ret = s.data[s.sz]
	s.data = s.data[:s.sz]
	return
}

func (s *stack) Empty() bool {
	if s.sz == 0 {
		return true
	}
	return false
}

func Constructor() CQueue {
	return CQueue{}
}

func (ccc *CQueue) AppendTail(value int) {
	s1 := &(ccc.s1)
	s1.Push(value)
}

func (ccc *CQueue) DeleteHead() int {
	s1 := &(ccc.s1)
	s2 := &(ccc.s2)
	if s2.Empty() {
		if s1.Empty() {
			return -1
		} else {
			for s1.Empty() == false {
				t := s1.Pop()
				s2.Push(t)
			}
			return s2.Pop()
		}
	} else {
		return s2.Pop()
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
