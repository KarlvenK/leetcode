package main

import "fmt"

func numberOfWays(s string) int64 {
	buildings := []byte(s)
	n := len(buildings)

	num0, num1 := make([]int64, n), make([]int64, n)
	if buildings[0] == '0' {
		num0[0] = 1
		num1[0] = 0
	} else {
		num0[0] = 0
		num1[0] = 1
	}
	for i := 1; i < n; i++ {
		if buildings[i] == '0' {
			num0[i] = num0[i-1] + 1
			num1[i] = num1[i-1]
		} else {
			num0[i] = num0[i-1]
			num1[i] = num1[i-1] + 1
		}
	}
	var ans int64
	for i := 1; i < n-1; i++ {
		if buildings[i] == '0' {
			ans += num1[i-1] * (num1[n-1] - num1[i])
		} else {
			ans += num0[i-1] * (num0[n-1] - num0[i])
		}
	}

	return ans
}

func main() {
	s := "001101"
	fmt.Println(numberOfWays(s))
}
