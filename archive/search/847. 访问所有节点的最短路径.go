package search

import (
	"container/list"
	"fmt"
)

type node struct {
	head    int
	covered int
}

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	que := list.New()
	dest := (1 << n) - 1
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, 1<<n)
		for j := 0; j < (1 << n); j++ {
			dist[i][j] = n * n
		}
	}

	for i := 0; i < n; i++ {
		que.PushBack(node{i, 1 << i})
		dist[i][1<<i] = 0
	}

	for que.Len() != 0 {
		u := que.Front().Value.(node)
		que.Remove(que.Front())
		if u.covered == dest {
			return dist[u.head][u.covered]
		}

		for _, v := range graph[u.head] {
			_covered_ := u.covered | (1 << v)
			if dist[u.head][u.covered]+1 < dist[v][_covered_] {
				dist[v][_covered_] = dist[u.head][u.covered] + 1
				que.PushBack(node{v, _covered_})
			}
		}
	}
	return -1
}

func main() {
	var graph [][]int
	for i := 0; i < 4; i++ {
		graph = append(graph, make([]int, 0))
		graph[i] = append(graph[i], 0)
	}
	graph[0] = nil
	graph[0] = append(graph[0], 1, 2, 3)
	fmt.Println(shortestPathLength(graph))

}

/*

给出 graph 为有 N 个节点（编号为 0, 1, 2, ..., N-1）的无向连通图。

graph.length = N，且只有节点 i 和 j 连通时，j != i 在列表 graph[i] 中恰好出现一次。

返回能够访问所有节点的最短路径的长度。你可以在任一节点开始和停止，也可以多次重访节点，并且可以重用边。

solution:
以“已经访问过的点”为状态，用位运算压缩状态(covered)。covered的二进制形式的第i位为1时，表示已经访问过i了。
然后用bfs算法。dist[x][covered_x] 表示访问到x时，covered = covered_x时的最短路径。
当我们从que提取的node的covered为11111...111时，则说明已经完全访问。直接返回答案。
*/
