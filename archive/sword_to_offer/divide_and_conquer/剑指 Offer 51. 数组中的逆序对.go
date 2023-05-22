package divide_and_conquer

func reversePairs(nums []int) int {
	var dfs func(int, int)
	tmp := make([]int, len(nums))
	ans := 0
	dfs = func(left, right int) {
		if left >= right {
			return
		}
		mid := (left + right) >> 1
		dfs(left, mid)
		dfs(mid+1, right)
		i, j := left, mid+1
		pos := 0
		for i <= mid && j <= right {
			if nums[i] <= nums[j] {
				tmp[pos] = nums[i]
				pos++
				i++
				ans += j - mid - 1
			} else {
				tmp[pos] = nums[j]
				pos++
				j++
			}
		}
		for i <= mid {
			tmp[pos] = nums[i]
			pos++
			ans += right - mid
			i++
		}
		for j <= right {
			tmp[pos] = nums[j]
			pos++
			j++
		}
		for t := left; t <= right; t++ {
			nums[t] = tmp[t-left]
		}
	}
	dfs(0, len(nums)-1)
	return ans
}
