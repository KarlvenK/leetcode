package main

import "fmt"

func countBits(num int) []int {
	bits := make([]int, num+1)
	for i := 1; i <= num; i++ {
		bits[i] = bits[i>>1] + i&1
	}
	return bits
}

func main() {
	num := 10
	ans := countBits(num)
	for i := 0; i <= num; i++ {
		fmt.Printf("%d ", ans[i])
	}
}
