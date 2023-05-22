package main

import "fmt"

func minBitFlips(start int, goal int) int {
	//a, b := start, goal
	bit1 := make([]int, 0)
	for start != 0 {
		bit1 = append(bit1, start&1)
		start = start >> 1
	}
	bit2 := make([]int, 0)
	for goal != 0 {
		bit2 = append(bit2, goal&1)
		goal = goal >> 1
	}
	l := max(len(bit1), len(bit2))
	for len(bit1) != l {
		bit1 = append(bit1, 0)
	}
	for len(bit2) != l {
		bit2 = append(bit2, 0)
	}
	ans := 0
	for i := 0; i < l; i++ {
		if bit1[i] != bit2[i] {
			ans++
		}
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(10, 20)
}
