package search_easy

import "container/list"

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	lis := list.New()
	lis.PushBack(root)
	temp := list.New()
	for lis.Len() != 0 {
		tool := make([]int, 0)
		for lis.Len() != 0 {
			e := lis.Front()
			node := e.Value.(*TreeNode)
			lis.Remove(e)
			if node.Left != nil {
				temp.PushBack(node.Left)
			}
			if node.Right != nil {
				temp.PushBack(node.Right)
			}
			tool = append(tool, node.Val)
		}
		for temp.Len() != 0 {
			e := temp.Front()
			lis.PushBack(e.Value.(*TreeNode))
			temp.Remove(e)
		}
		ans = append(ans, tool)
	}
	return
}
