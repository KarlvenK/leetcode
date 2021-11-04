package two_ptr

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Val: -1,
	}
	dummy.Next = head
	cur := head
	pre := dummy
	for cur != nil {
		if cur.Val == val {
			tmp := cur.Next
			pre.Next = tmp
			break
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return dummy.Next
}
