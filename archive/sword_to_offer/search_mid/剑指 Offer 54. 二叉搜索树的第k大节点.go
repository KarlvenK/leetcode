package search_mid

func kthLargest(root *TreeNode, k int) int {
	var dfs func(*TreeNode)
	cnt := 0
	ans := -1
	dfs = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		dfs(cur.Right)
		cnt++
		if cnt == k {
			ans = cur.Val
		}
		dfs(cur.Left)
	}
	dfs(root)
	return ans
}
