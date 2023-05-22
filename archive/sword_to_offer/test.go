package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 5, 5, 6, 6, 7, 7, 8, 9}
	var search func(int) int
	search = func(x int) int {
		l, r := 0, len(nums)
		for l != r {
			mid := (l + r) >> 1
			if nums[mid] < x {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return r
	}
	for i := 1; i < 20; i++ {
		fmt.Println(search(i))
	}
}
