package main

func allPossibleFBT(n int) (ans []*TreeNode) {
	if n == 1 {
		ans = append(ans, &TreeNode{
			Val: 0,
		})
		return
	}

	for i := 1; i <= n-2; i += 2 {
		leftAns := allPossibleFBT(i)
		rightAns := allPossibleFBT(n - i - 1)
		for j := 0; j < len(leftAns); j++ {
			for k := 0; k < len(rightAns); k++ {
				temp := &TreeNode{
					Val: 0,
				}
				temp.Left = leftAns[j]
				temp.Right = rightAns[k]
				ans = append(ans, temp)
			}
		}
	}
	return
}

/*
满二叉树是一类二叉树，其中每个结点恰好有 0 或 2 个子结点。
返回包含 N 个结点的所有可能满二叉树的列表。 答案的每个元素都是一个可能树的根结点。
答案中每个树的每个结点都必须有 node.val=0。
你可以按任何顺序返回树的最终列表。


n = 1 [0]
n > 1 时，根据左右字数节点个数分类

采用分治，最后将左右字数的所有情况merge
*/
