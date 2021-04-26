package binary

func shipWithinDays(weights []int, D int) int {
	var maxInt func(int, int) int
	maxInt = func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	left, right := 0, 0
	//n := len(weights)
	for _, v := range weights {
		right += v
		left = maxInt(left, v)
	}

	var check func(int) bool
	check = func(capacity int) bool {
		temp := 0
		cnt := 0
		for _, v := range weights {
			if temp+v <= capacity {
				temp += v
			} else {
				temp = v
				cnt++
			}
		}
		return cnt <= D
	}

	for left != right {
		mid := left + ((right - left) >> 1)
		if check(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
