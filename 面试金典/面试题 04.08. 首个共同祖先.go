package interview

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	var dfs func (*TreeNode) *TreeNode

	dfs = func(cur *TreeNode) *TreeNode {
		if cur == nil {
			return nil
		}
		if cur == p || cur == q {
			return cur
		}
		l, r := dfs(cur.Left), dfs(cur.Right)
		if l != nil && r != nil {
			return cur
		}
		if l != nil {
			return l
		}
		if r != nil {
			return r
		}
		return nil
	}

	return dfs(root)
}