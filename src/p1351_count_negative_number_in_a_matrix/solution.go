package p1351_count_negative_number_in_a_matrix

func countNegatives(grid [][]int) int {
	ans := 0
	lastPos := len(grid[0])
	for _, row := range grid {
		l, r := 0, lastPos
		pos := lastPos
		for l < r {
			mid := (r-l)>>1 + l
			if row[mid] >= 0 {
				l = mid + 1
			} else {
				pos = mid
				r = mid
			}
		}
		ans += len(row) - pos
	}
	return ans
}

/*
 4  3  2 -1
 3  2  1 -1
 1  1 -1 -2
-1 -1 -2 -3
*/
