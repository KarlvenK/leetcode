package main

import "fmt"

type Stack struct {
	data []int
}

func (s *Stack) Pop() int {
	num := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return num
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *Stack) Top() int {
	return s.data[len(s.data)-1]
}

func New() Stack {
	t := new(Stack)
	t.data = make([]int, 0)
	return *t
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func trap(height []int) int {
	n := len(height)
	s := New()
	ans := 0
	for i := 0; i < n; i++ {
		for s.Empty() == false && height[i] > height[s.Top()] {
			top := height[s.Pop()]
			if s.Empty() == true {
				break
			}
			left := s.Top()
			wid := i - left - 1
			hei := min(height[left], height[i]) - top
			ans += wid * hei
		}
		s.Push(i)
	}
	return ans
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
