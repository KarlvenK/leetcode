package divide_and_conquer

func verifyPostorder(postorder []int) bool {
	n := len(postorder)
	var dfs func(int, int) bool
	dfs = func(left, right int) bool {
		if left >= right {
			return true
		}
		p := left
		for postorder[p] < postorder[right] {
			p++
		}
		mid := p
		for postorder[p] > postorder[right] {
			p++
		}
		if p != right {
			return false
		}
		return dfs(left, mid-1) && dfs(mid, right-1)
	}
	return dfs(0, n-1)
}
