package graph

type edge struct {
	v int
	w int
}

func networkDelayTime(times [][]int, n int, k int) int {
	used := make([]bool, n+1)
	graph := make([][]edge, n+1)
	ans := -1
	for i := 0; i < len(times); i++ {
		u, v, w := times[i][0], times[i][1], times[i][2]
		graph[u] = append(graph[u], edge{v, w})
	}
	var que = []int{k}

	dist := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dist[i] = 1234567890
	}
	dist[k] = 0

	for len(que) != 0 {
		u := que[0]
		que = que[1:]
		used[u] = false
		for i := 0; i < len(graph[u]); i++ {
			v, w := graph[u][i].v, graph[u][i].w
			if dist[v] > (dist[u] + w) {
				dist[v] = dist[u] + w
				if !used[v] {
					que = append(que, v)
					used[v] = true
				}
			}
		}
	}
	for i := 0; i <= n; i++ {
		ans = max(ans, dist[i])
	}
	if ans == 1234567890 {
		return -1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
