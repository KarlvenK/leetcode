package binary

func Search(nums []int, target int) int {
	n := len(nums)
	if n < 1 {
		return -1
	}
	if n == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	left, right := 0, n
	//[left, right)
	for left+1 < right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[n-1] {
				left = mid + 1
			} else {
				right = mid
			}
		}

	}
	if left >= n || left < 0 {
		return -1
	}
	if nums[left] == target {
		return left
	}
	return -1
}
