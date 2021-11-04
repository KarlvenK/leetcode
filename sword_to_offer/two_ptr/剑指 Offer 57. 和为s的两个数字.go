package two_ptr

func twoSum(nums []int, target int) (ans []int) {
	p1, p2 := 0, len(nums)-1

	for true {
		if nums[p1]+nums[p2] == target {
			ans = append(ans, nums[p1], nums[p2])
			break
		}
		if nums[p1]+nums[p2] > target {
			p2--
		} else {
			p1++
		}
	}
	return
}
