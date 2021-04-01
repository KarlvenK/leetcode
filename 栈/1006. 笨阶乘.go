package main

func clumsy(n int) int {
	stack := []int{n}
	n--
	t := 0
	for n > 0 {
		switch t % 4 {
		case 0:
			stack[len(stack)-1] *= n
		case 1:
			stack[len(stack)-1] /= n
		case 2:
			stack = append(stack, n)
		case 3:
			stack = append(stack, -n)
		}
		n--
		t++
	}
	ans := 0
	for _, v := range stack {
		ans += v
	}
	return ans
}
