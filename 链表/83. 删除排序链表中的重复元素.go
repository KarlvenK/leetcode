package main

func deleteDuplicates(head *ListNode) *ListNode {
	/*
		dummyHead := &ListNode{
			-1,
			nil,
		}
		pre := dummyHead
		cur := head
		mark := -123456
		for cur != nil {
			if cur.Val != mark {
				mark = cur.Val
				pre.Next = cur
				pre = cur
				cur = cur.Next
				pre.Next = nil
			} else {
				cur = cur.Next
			}
		}
		return dummyHead.Next
	*/
	// better version
	if head == nil {
		return nil
	}
	cur := head
	for cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {

}
