package main

func stoneGameIX(stones []int) bool {
	a, b, c := 0, 0, 0
	for _, v := range stones {
		mod := v % 3
		if mod == 0 {
			a++
		}
		if mod == 1 {
			b++
		}
		if mod == 2 {
			c++
		}
	}

	if a%2 == 0 {
		return b*c > 0
	} else {
		return abs(b-c) >= 3
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}
