package search_easy

func mirrorTree(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		node.Left, node.Right = node.Right, node.Left
		node.Left = dfs(node.Left)
		node.Right = dfs(node.Right)
		return node
	}
	return dfs(root)
}
