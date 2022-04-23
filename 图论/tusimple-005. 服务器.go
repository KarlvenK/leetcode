package main

import "fmt"

type edge struct {
	to   int
	dist int
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	nameToInt := make(map[string]int)
	cnt := 0
	graph := make([][]edge, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]edge, 0)
	}

	for i := 0; i < m; i++ {
		var a, b string
		var dis int
		fmt.Scan(&a, &b, &dis)
		if _, ok := nameToInt[a]; !ok {
			nameToInt[a] = cnt
			cnt++
		}
		if _, ok := nameToInt[b]; !ok {
			nameToInt[b] = cnt
			cnt++
		}
		u, v := nameToInt[a], nameToInt[b]
		graph[u] = append(graph[u], edge{
			to:   v,
			dist: dis,
		})
		// graph[v] = append(graph[v], edge {
		//     to: u,
		//     dist: dis,
		// })
	}

	var q int
	fmt.Scan(&q)

	dijkstra := func(u int, v int) int {
		dis := make([]int, n)
		mark := make([]bool, n)
		for i := 0; i < n; i++ {
			dis[i] = 12345678901
		}
		dis[u] = 0

		for t := 0; t < n; t++ {
			min := 12345678901
			k := -1
			for i := 0; i < n; i++ {
				if mark[i] == false && dis[i] < min {
					min = dis[i]
					k = i
				}
			}
			if k == -1 {
				break
			}
			mark[k] = true

			for j := 0; j < len(graph[k]); j++ {
				edge := graph[k][j]
				v := edge.to
				if mark[v] == false {
					if dis[v] > dis[k]+edge.dist {
						dis[v] = dis[k] + edge.dist
					}
				}
			}
		}
		if mark[v] == false {
			return -1
		}
		return dis[v]
	}

	for t := 0; t < q; t++ {
		var a, b string
		fmt.Scan(&a, &b)
		u, v := nameToInt[a], nameToInt[b]
		ans := dijkstra(u, v)
		if ans == -1 {
			fmt.Println("INF")
		} else {
			fmt.Println(ans)
		}
	}

}
