package main

func xorQueries(arr []int, queries [][]int) (ans []int) {

	sum := make([]int, len(arr)+1)
	sum[0] = 0
	for i, e := range arr {
		sum[i+1] = sum[i] ^ e
	}
	// sum[i] 表示 [0 .. i - 1]的亦或和
	for i := range queries {
		ans = append(ans, sum[queries[i][0]]^sum[queries[i][1]+1])
	}
	return
}
