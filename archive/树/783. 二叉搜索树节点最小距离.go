package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	ans := math.MaxInt32
	pre := -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 {
			ans = min(ans, node.Val-pre)
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
