package p1123_lca_of_deepest_leaves

import . "github.com/KarlvenK/leetcode/src/utils"

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, ans := dfs(root, 0)
	return ans
}

func dfs(cur *TreeNode, depth int) (int, *TreeNode) {
	if cur == nil {
		return -1, nil
	}
	if cur.Left == nil && cur.Right == nil {
		return depth, cur
	}
	leftDepth, leftLCA := dfs(cur.Left, depth+1)
	rightDepth, rightLCA := dfs(cur.Right, depth+1)
	if leftLCA == nil {
		return rightDepth, rightLCA
	}
	if rightLCA == nil {
		return leftDepth, leftLCA
	}
	if leftDepth == rightDepth {
		return leftDepth, cur
	}
	if leftDepth > rightDepth {
		return leftDepth, leftLCA
	}
	return rightDepth, rightLCA
}
