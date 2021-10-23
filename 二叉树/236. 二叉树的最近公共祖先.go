package main

// TreeNode declared
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent := make(map[*TreeNode]*TreeNode)
	var dfs func(*TreeNode)
	dfs = func(cur *TreeNode) {
		if cur.Left != nil {
			parent[cur.Left] = cur
			dfs(cur.Left)
		}
		if cur.Right != nil {
			parent[cur.Right] = cur
			dfs(cur.Right)
		}
	}
	dfs(root)
	used := make(map[*TreeNode]struct{})

	t := p
	for true {
		used[t] = struct{}{}
		if _, ok := parent[t]; !ok {
			break
		}
		t = parent[t]
	}
	t = q
	for true {
		if _, ok := used[t]; ok {
			return t
		}
		t = parent[t]
	}
	return nil
}
