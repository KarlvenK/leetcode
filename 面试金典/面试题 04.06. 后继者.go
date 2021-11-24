package interview

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var pre *TreeNode
	for root != p {
		if root.Val < p.Val {
			root = root.Right
		} else {
			pre = root
			root = root.Left
		}
	}

	if root.Right == nil {
		root = pre
	} else {
		root = root.Right
		for root.Left != nil {
			root = root.Left
		}
	}
	return root
}
