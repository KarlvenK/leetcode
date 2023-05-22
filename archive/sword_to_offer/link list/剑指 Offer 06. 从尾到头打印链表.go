package link_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) (ans []int) {
	var dfs func(node *ListNode)
	dfs = func(node *ListNode) {
		if node == nil {
			return
		}
		dfs(node.Next)
		ans = append(ans, node.Val)
	}
	dfs(head)
	return
}
