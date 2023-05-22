package search

import "sort"

func SubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var dfs func(bool, int)
	ans := make([][]int, 0)
	temp := make([]int, 0)
	dfs = func(choosePre bool, idx int) {
		if idx == n {
			//ans = append(ans, append([]int(nil), temp...))
			ans = append(ans, append(make([]int, 0), temp...))
			/*var auto []int
			for _, v := range temp {
				auto = append(auto, v)
			}
			ans = append(ans, auto)*/
			return
		}
		dfs(false, idx+1)
		if choosePre == false && idx > 0 && nums[idx] == nums[idx-1] {
			return
		}
		temp = append(temp, nums[idx])
		dfs(true, idx+1)
		temp = temp[:len(temp)-1]
	}
	dfs(false, 0)
	return ans
}
