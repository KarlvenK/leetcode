package search_mid

func exist(board [][]byte, word string) bool {
	n, m := len(board), len(board[0])
	var dfs func(int, int, int) bool
	check := func(x, y int) bool {
		if x < 0 || y < 0 {
			return false
		}
		if x >= n || y >= m {
			return false
		}
		if board[x][y] == ' ' {
			return false
		}
		return true
	}

	var delta = [4][2]int{
		{1, 0}, {0, 1},
		{-1, 0}, {0, -1},
	}

	dfs = func(x, y, step int) bool {
		if board[x][y] != word[step] {
			return false
		}
		if step == len(word)-1 {
			return true
		}
		var dx, dy int
		tmp := board[x][y]
		board[x][y] = ' '
		for i := 0; i < 4; i++ {
			dx, dy = x+delta[i][0], y+delta[i][1]
			if check(dx, dy) {
				if dfs(dx, dy, step+1) {
					return true
				}
			}
		}
		board[x][y] = tmp
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
