package main

import "fmt"

type unionSet struct {
	parent []int
}

func (s unionSet) find(x int) int {
	if x != s.parent[x] {
		s.parent[x] = s.find(s.parent[x])
	}
	return s.parent[x]
}

func (s unionSet) merge(a, b int) {
	pa, pb := s.find(a), s.find(b)
	if pa != pb {
		s.parent[pa] = s.parent[pb]
	}
}

func newUnion(n int) unionSet {
	var ret unionSet
	ret.parent = make([]int, n)
	for i := 0; i < n; i++ {
		ret.parent[i] = i
	}
	return ret
}

func largestComponentSize(nums []int) int {
	maxN := 0
	for _, x := range nums {
		maxN = max(x, maxN)
	}
	us := newUnion(maxN + 1)
	for _, num := range nums {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				us.merge(num, i)
				us.merge(num, num/i)
			}
		}
	}
	cnt := make([]int, maxN+1)
	ans := 0
	for _, num := range nums {
		cnt[us.find(num)]++
		ans = max(ans, cnt[us.find(num)])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func _gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return _gcd(b, a%b)
}

func main() {
	arr := [][]int{
		{4, 6, 15, 35},
		{20, 50, 9, 63},
		{2, 3, 6, 7, 4, 12, 21, 39},
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(largestComponentSize(arr[i]))
	}
}
