package divide_and_conquer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	tool := make(map[int]int)
	n := len(preorder)
	for i := 0; i < n; i++ {
		tool[inorder[i]] = i
	}
	var dfs func(pl, pr, il, ir int) *TreeNode
	dfs = func(pl, pr, il, ir int) *TreeNode {
		if pl > pr {
			return nil
		}
		root := &TreeNode{
			Val:   preorder[pl],
			Left:  nil,
			Right: nil,
		}
		pos := tool[preorder[pl]]
		length := pos - il
		root.Left = dfs(pl+1, pl+length, il, pos-1)
		root.Right = dfs(pl+length+1, pr, pos+1, ir)
		return root
	}
	return dfs(0, n-1, 0, n-1)
}
