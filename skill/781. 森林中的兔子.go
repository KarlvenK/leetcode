package main

import "fmt"

func numRabbits(answers []int) int {
	tool := make(map[int]int)
	var ans = 0
	for _, v := range answers {
		tool[v]++
	}
	for key, val := range tool {
		ans += (val / (key + 1)) * (key + 1)
		if val%(key+1) != 0 {
			ans += key + 1
		}
	}
	return ans
}

func main() {
	answer := []int{1, 1, 2}
	fmt.Println(numRabbits(answer))
}
