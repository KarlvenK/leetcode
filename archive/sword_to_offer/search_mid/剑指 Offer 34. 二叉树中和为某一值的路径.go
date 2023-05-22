package search_mid

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) (ans [][]int) {
	path := make([]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(cur *TreeNode, sum int) {
		if cur == nil {
			return
		}
		path = append(path, cur.Val)
		defer func() { // 如果此处不使用defer，那么！！！！！处的return不应该加。因为我们必须保证在退出一层dfs时，path末尾元素要弹出
			path = path[:len(path)-1]
		}()
		sum += cur.Val
		//fmt.Println(sum, path)
		if sum == target && cur.Left == nil && cur.Right == nil {
			t := make([]int, 0)
			ans = append(ans, append(t, path...))
			return //！！！！！
		}
		dfs(cur.Left, sum)
		dfs(cur.Right, sum)
	}
	dfs(root, 0)
	return
}
