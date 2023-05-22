package search_easy

func isSymmetric(root *TreeNode) bool {
	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}

	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}

/*
dfs在递归中比较两颗子树left和right
然后left的左子树与right的右子树比较
递归下去
*/
