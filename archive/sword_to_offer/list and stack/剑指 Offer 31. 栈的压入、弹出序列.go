package list_and_stack

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	stk := make([]int, n)
	top := 0
	idx := 0
	for _, num := range pushed {
		stk[top] = num
		top++
		for stk[top-1] == popped[idx] {
			top--
			idx++
			if top == 0 || idx == n {
				break
			}
		}
	}
	return top == 0
}
