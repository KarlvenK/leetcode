package two_ptr

func exchange(nums []int) []int {
	n := len(nums)
	p1, p2 := 0, n-1
	for p1 < p2 {
		for p1 < p2 && nums[p1]%2 == 1 {
			p1++
		}
		for p1 < p2 && nums[p2]%2 == 0 {
			p2--
		}
		if p1 < p2 {
			nums[p1], nums[p2] = nums[p2], nums[p1]
		}
	}
	return nums
}
