package search_easy

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (ans []int) {
	lis := list.New()
	lis.PushBack(root)
	if root == nil {
		return
	}
	for lis.Len() != 0 {
		e := lis.Front()
		lis.Remove(e)
		node := e.Value.(*TreeNode)
		ans = append(ans, node.Val)
		if node.Left != nil {
			lis.PushBack(node.Left)
		}
		if node.Right != nil {
			lis.PushBack(node.Right)
		}
	}
	return
}
