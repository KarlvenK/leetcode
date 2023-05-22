package interview

type TreeNode struct {
	Left 	*TreeNode
	Right 	*TreeNode
	Val 	int
}

type ListNode struct {
	Val int
	Next *ListNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}