package twoPtr

func removeDuplicates(nums []int) int {
	p1, p2 := 2, 2
	n := len(nums)
	for p2 < n {
		if nums[p1-2] != nums[p2] {
			nums[p1] = nums[p2]
			p1++
		}
		p2++
	}
	return p1
}
