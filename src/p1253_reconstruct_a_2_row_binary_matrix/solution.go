package p1253_reconstruct_a_2_row_binary_matrix

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)
	ans := make([][]int, 2)
	ans[0], ans[1] = make([]int, n), make([]int, n)

	sum := 0
	cnt := 0
	for i := 0; i < n; i++ {
		sum += colsum[i]
		if colsum[i] == 2 {
			cnt++
		}
	}
	if sum != upper+lower {
		return [][]int{}
	}
	if cnt > min(upper, lower) {
		return [][]int{}
	}
	upper -= cnt
	lower -= cnt

	for i := 0; i < n; i++ {
		if colsum[i] == 2 {
			ans[0][i], ans[1][i] = 1, 1
		}
		if colsum[i] == 1 {
			if upper > 0 {
				ans[0][i] = 1
				upper--
			} else {
				ans[1][i] = 1
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
