package main

func deleteDuplicates(head *ListNode) *ListNode {
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
}

func main() {

}
