package search

func gridIllumination(n int, lamps [][]int, queries [][]int) []int {
	row, column, leftToRight, rightToLeft := map[int]int{}, map[int]int{}, map[int]int{}, map[int]int{}

	type light struct {
		x int
		y int
	}

	hasLight := make(map[light]struct{})

	for _, loc := range lamps {
		if _, ok := hasLight[light{
			x: loc[0],
			y: loc[1],
		}]; ok {
			continue
		}
		hasLight[light{
			x: loc[0],
			y: loc[1],
		}] = struct{}{}
		row[loc[0]]++
		column[loc[1]]++
		leftToRight[n-loc[1]+loc[0]]++
		rightToLeft[loc[0]+loc[1]]++
	}

	ans := make([]int, len(queries))

	delta := [9][2]int{
		{0, 0},
		{1, 0}, {0, 1}, {-1, 0}, {0, -1},
		{1, 1}, {-1, -1}, {1, -1}, {-1, 1},
	}

	for i := 0; i < len(queries); i++ {
		if row[queries[i][0]] > 0 || column[queries[i][1]] > 0 || leftToRight[n-queries[i][1]+queries[i][0]] > 0 || rightToLeft[queries[i][0]+queries[i][1]] > 0 {
			ans[i] = 1
			for j := 0; j < 9; j++ {
				dx, dy := queries[i][0]+delta[j][0], queries[i][1]+delta[j][1]
				if dx < 0 || dx >= n || dy < 0 || dy >= n {
					continue
				}
				if _, ok := hasLight[light{
					x: dx,
					y: dy,
				}]; ok {
					delete(hasLight, light{
						x: dx,
						y: dy,
					})
					row[dx]--
					column[dy]--
					leftToRight[n-dy+dx]--
					rightToLeft[dx+dy]--
				}
			}

		} else {
			ans[i] = 0
		}
	}

	return ans
}
