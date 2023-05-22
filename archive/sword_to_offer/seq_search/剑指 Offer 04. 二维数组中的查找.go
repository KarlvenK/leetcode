package seq_search

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	dx, dy := n-1, 0

	check := func(x, y int) bool {
		if x < 0 || x >= n {
			return false
		}
		if y < 0 || y >= m {
			return false
		}
		return true
	}

	for check(dx, dy) {
		if matrix[dx][dy] == target {
			return true
		}
		if matrix[dx][dy] > target {
			dx = dx - 1
		} else {
			dy = dy + 1
		}
	}
	return false
}
