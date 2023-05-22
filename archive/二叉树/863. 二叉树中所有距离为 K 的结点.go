package main

// TreeNode declared
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) (ans []int) {
	parent := make(map[int]*TreeNode)
	used := make(map[*TreeNode]struct{})

	var getParent func(*TreeNode, *TreeNode)
	getParent = func(cur, p *TreeNode) {
		parent[cur.Val] = p
		if cur.Left != nil {
			getParent(cur.Left, cur)
		}
		if cur.Right != nil {
			getParent(cur.Right, cur)
		}
	}
	getParent(root, root)

	var dfs func(*TreeNode, int)
	dfs = func(cur *TreeNode, depth int) {
		if depth == k {
			ans = append(ans, cur.Val)
			return
		}
		used[cur] = struct{}{}
		if cur.Left != nil {
			if _, ok := used[cur.Left]; !ok {
				dfs(cur.Left, depth+1)
			}
		}
		if cur.Right != nil {
			if _, ok := used[cur.Right]; !ok {
				dfs(cur.Right, depth+1)
			}
		}
		if _, ok := used[parent[cur.Val]]; !ok {
			dfs(parent[cur.Val], depth+1)
		}
	}

	dfs(target, 0)
	return
}
