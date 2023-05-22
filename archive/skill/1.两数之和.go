package main

import "fmt"

func twoSum(nums []int, target int) []int {
	table := make(map[int][]int)
	for i, num := range nums {
		table[num] = append(table[num], i)
	}
	for i, num := range nums {
		temp, ok := table[target-num]
		if ok {
			if len(temp) == 1 && temp[0] == i {
				continue
			} else {
				ans := make([]int, 2)
				ans[0] = i
				for j := 0; j < len(temp); j++ {
					if temp[j] != i {
						ans[1] = temp[j]
						return ans
					}
				}
			}
		}
	}
	return make([]int, 0)
}

func main() {
	nums := []int{2, 2, 7, 11, 15}
	fmt.Println(twoSum(nums, 4))
}
