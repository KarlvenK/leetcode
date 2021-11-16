package sorting

import (
	"fmt"
	"sort"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x := fmt.Sprintf("%d%d", nums[i], nums[j])
		y := fmt.Sprintf("%d%d", nums[j], nums[i])
		return x < y
	})
	var ans string
	for i := 0; i < len(nums); i++ {
		ans += fmt.Sprintf("%d", nums[i])
	}
	return ans
}
