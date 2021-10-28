package search_easy

import "container/list"

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	leftToRight := true
	queue := list.New()
	tmp := list.New()
	queue.PushBack(root)

	for queue.Len() != 0 {
		_ans := make([]int, 0)
		for queue.Len() != 0 {
			e := queue.Front()
			queue.Remove(e)
			node := e.Value.(*TreeNode)
			_ans = append(_ans, node.Val)
			if node.Left != nil {
				tmp.PushBack(node.Left)
			}
			if node.Right != nil {
				tmp.PushBack(node.Right)
			}
		}
		for tmp.Len() != 0 {
			e := tmp.Front()
			node := e.Value.(*TreeNode)
			tmp.Remove(e)
			queue.PushBack(node)
		}
		if !leftToRight {
			reverse(_ans)
		}
		leftToRight = !leftToRight
		ans = append(ans, _ans)
	}
	return
}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
