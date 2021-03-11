package main

import "fmt"

func Calculate(s string) int {
	opt := []int{1}
	sign := 1
	ans := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = opt[len(opt)-1]
			i++
		case '-':
			sign = -opt[len(opt)-1]
			i++
		case '(':
			opt = append(opt, sign)
			i++
		case ')':
			opt = opt[:len(opt)-1]
			i++
		default:
			num := 0
			for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
				num = num*10 + (int)(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return ans
}

func main() {
	s := "(1+(4+5+2)-3)+(6+8 )"
	fmt.Println(Calculate(s))
}
