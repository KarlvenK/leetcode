package search_mid

import "container/list"

func movingCount(m int, n int, k int) int {
	que := list.New()
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	que.PushBack([2]int{0, 0})

	check := func(x, y int) bool {
		if x < 0 || y < 0 || x >= m || y >= n {
			return false
		}
		cnt := 0
		for x != 0 {
			cnt += x % 10
			x /= 10
		}
		for y != 0 {
			cnt += y % 10
			y /= 10
		}
		if cnt > k {
			return false
		}
		return true
	}
	var delta = [4][2]int{
		{1, 0}, {0, 1},
		{0, 1}, {0, -1},
	}
	ans := 1
	vis[0][0] = true
	for que.Len() != 0 {
		pos := que.Front().Value.([2]int)
		que.Remove(que.Front())
		//fmt.Println(pos[0], pos[1])
		for i := 0; i < 2; i++ {
			dx := pos[0] + delta[i][0]
			dy := pos[1] + delta[i][1]
			if check(dx, dy) {
				if !vis[dx][dy] {
					que.PushBack([2]int{dx, dy})
					ans++
					vis[dx][dy] = true
				}
			}
		}
	}
	return ans
}
