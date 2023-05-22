package interview

func removeDuplicateNodes(head *ListNode) *ListNode {
	dict := make(map[int]struct{})
	dummy := &ListNode{
		Val:  -1,
		Next: head,
	}
	var cur, pre *ListNode
	pre = dummy
	cur = head
	for cur != nil {
		if _, ok := dict[cur.Val]; ok {
			pre.Next = cur.Next
			cur = cur.Next
		} else {
			dict[cur.Val] = struct{}{}
			pre = cur
			cur = cur.Next
		}
	}
	return dummy.Next
}
