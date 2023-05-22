package search

import "fmt"

func slidingPuzzle(board [][]int) int {
	const target = "123450"
	mark := map[string]bool{}
	var temp []byte
	var neighbors = [6][]int{
		{1, 3},
		{0, 2, 4},
		{1, 5},
		{0, 4},
		{1, 3, 5},
		{2, 4},
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			temp = append(temp, byte(board[i][j]+int('0')))
		}
	}
	start := string(temp)
	if start == target {
		return 0
	}

	mark[start] = true

	type pair struct {
		status string
		step   int
	}

	getNext := func(cur string) (ret []string) {
		temp := []byte(cur)
		pos := -1
		for i, b := range temp {
			if b == '0' {
				pos = i
				break
			}
		}

		for _, x := range neighbors[pos] {
			temp[pos], temp[x] = temp[x], temp[pos]
			ret = append(ret, string(temp))
			temp[pos], temp[x] = temp[x], temp[pos]
		}

		return
	}
	que := []pair{
		{
			start,
			0,
		},
	}

	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		for _, nx := range getNext(p.status) {
			if mark[nx] != true {
				if nx == target {
					return p.step + 1
				}
				mark[nx] = true
				que = append(que, pair{
					nx,
					p.step + 1,
				})
			}
		}
	}

	return -1
}

func main() {
	board := [][]int{{1, 2, 3}, {5, 4, 0}}
	fmt.Println(slidingPuzzle(board))
}
