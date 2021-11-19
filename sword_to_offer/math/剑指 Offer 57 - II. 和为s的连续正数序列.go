package math

func findContinuousSequence(target int) (ans [][]int) {
	n := target>>1 + 1
	preSum := make([]int, n+1)
	tool := make(map[int]int)
	tool[0] = 0
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + i
		tool[preSum[i]] = i
	}

	for i := 1; i <= n; i++ {
		delta := preSum[i] - target
		if delta >= 0 {
			if index, ok := tool[delta]; ok {
				tmp := make([]int, 0)
				for j := index + 1; j <= i; j++ {
					tmp = append(tmp, j)
				}
				ans = append(ans, tmp)
			}
		}
	} /*
	   sort.Slice(ans,func(i, j int) bool{
	       return ans[i][0] < ans[j][0]
	   })*/
	return ans

}
