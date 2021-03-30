package search

import (
	"fmt"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	nums := []int{1, 2, 2}
	ans := SubsetsWithDup(nums)
	fmt.Print(ans)
}
