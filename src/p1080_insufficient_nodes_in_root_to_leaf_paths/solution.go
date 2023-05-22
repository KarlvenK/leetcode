package p1080_insufficient_nodes_in_root_to_leaf_paths

import (
	. "github.com/KarlvenK/leetcode/src/utils"
)

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	var dfs func(node *TreeNode, sum int) bool
	dfs = func(node *TreeNode, sum int) bool {
		if node == nil {
			return false
		}
		if node.Left == nil && node.Right == nil {
			return node.Val+sum >= limit
		}
		if dfs(node.Left, node.Val+sum) == false {
			node.Left = nil
		}
		if dfs(node.Right, node.Val+sum) == false {
			node.Right = nil
		}
		if node.Left == nil && node.Right == nil {
			return false
		}
		return true
	}
	if dfs(root, 0) == false {
		return nil
	}
	return root
}
