package main

import (
	"fmt"
	"strings"
)

func calculate(s string) int {
	table := map[byte]int{
		'-': 1,
		'+': 1,
		'*': 2,
		'/': 2,
	}

	s = strings.Replace(s, "(-", "(0-", -1)
	s = strings.Replace(s, " ", "", -1)
	n := len(s)

	nums := make([]int, 0)
	nums = append(nums, 0)
	ops := make([]byte, 0)

	for i := 0; i < n; i++ {
		ch := s[i]
		switch ch {
		case '(':
			ops = append(ops, ch)
		case ')':
			for len(ops) != 0 {
				if ops[len(ops)-1] != '(' {
					calc(&nums, &ops)
				} else {
					ops = ops[:len(ops)-1]
				}
			}
		default:
			if isNumber(ch) {
				u := 0
				j := i
				for ; j < n && isNumber(s[j]); j++ {
					u = u*10 + int(s[j]-'0')
				}
				nums = append(nums, u)
				i = j - 1
			} else {
				for len(ops) != 0 && ops[len(ops)-1] != '(' {
					prev := ops[len(ops)-1]
					if table[prev] >= table[ch] {
						calc(&nums, &ops)
					} else {
						break
					}
				}
				ops = append(ops, ch)
			}
		}
	}
	for len(ops) != 0 && ops[len(ops)-1] != '(' {
		calc(&nums, &ops)
	}
	return nums[len(nums)-1]
}

func isNumber(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func calc(nums *[]int, ops *[]byte) {
	if len(*nums) < 2 {
		return
	}
	if len(*ops) == 0 {
		return
	}

	b := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	a := (*nums)[len(*nums)-1]
	*nums = (*nums)[:len(*nums)-1]
	op := (*ops)[len(*ops)-1]
	*ops = (*ops)[:len(*ops)-1]

	res := 0
	switch op {
	case '+':
		res = a + b
	case '-':
		res = a - b
	case '*':
		res = a * b
	case '/':
		res = a / b
	}
	*nums = append(*nums, res)
}

func main() {
	s := " (13+5) / 2 "
	fmt.Println(calculate(s))
}
