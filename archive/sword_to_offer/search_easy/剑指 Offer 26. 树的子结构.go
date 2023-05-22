package search_easy

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	var dfs func(a, b *TreeNode) bool
	dfs = func(a, b *TreeNode) bool {
		if b == nil {
			return true
		}
		if a == nil && b != nil {
			return false
		}
		if a.Val == b.Val {
			if dfs(a.Left, b.Left) && dfs(a.Right, b.Right) {
				return true
			}
		}
		return dfs(a.Left, B) || dfs(a.Right, B)
	}
	if B == nil {
		return false
	}
	return dfs(A, B)
}

/*
整个树记为A和B
子树记为a和b

在dfs中使用a，b表示当前子树
如果当前b为nil， 则说明匹配完成，不管a是不是nil都是成立的
如果当前b不为nil
	如果a为nil， 则必然匹配失败
	如果a根值等于b根值
		那么匹配成功当且仅当a的左子树与b的左子树匹配，a的右子树和b的右子树匹配
		如果不相等
			那么就要让a的左右子树分别和B匹配，只要有一个匹配成功就算成功
*/
