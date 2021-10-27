package link_list

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	pre = nil

	for head != nil {
		t := head.Next
		head.Next = pre
		pre = head
		head = t
	}
	return pre
}
