package search

import "fmt"

func openLock(deadends []string, target string) int {
	const start = "0000"
	if start == target {
		return 0
	}
	dead := map[string]bool{}
	for _, t := range deadends {
		dead[t] = true
	}
	if dead[start] {
		return -1
	}

	type pair struct {
		nums string
		step int
	}

	que := []pair{
		{start,
			  0,
		},
	}
	mark := map[string]bool{
		start: true,
	}

	getNext := func(nums string) (ret []string) {
		str := []byte(nums)
		for i, char := range str {
			str[i] = char - 1
			if str[i] < '0' {
				str[i] = '9'
			}
			ret = append(ret, string(str))

			str[i] = char + 1
			if str[i] > '9' {
				str[i] = '0'
			}
			ret = append(ret, string(str))
			/*
			* 回溯， 枚举下一位
			*/
			str[i] = char
		}
		return
	}

	for len(que) > 0{
		node := que[0]
		que = que[1:]
		for _, v := range getNext(node.nums) {
			if v == target {
				return node.step + 1
			}
			if !mark[v] && !dead[v] {
				mark[v] =true
				que = append(que, pair{v, node.step + 1})
			}
		}
	}
	return -1
}

func main() {
	deadends := []string{"0201","0101","0102","1212","2002"}
	target := "0202"
	fmt.Println(openLock(deadends, target))
}
