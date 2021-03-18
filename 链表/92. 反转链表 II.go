package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	dummyHead.Next = head
	pre := dummyHead
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	rNode := pre
	for i := 0; i < right-left+1; i++ {
		rNode = rNode.Next
	}
	lNode := pre.Next
	succ := rNode.Next

	pre.Next = nil
	rNode.Next = nil
	reverse(lNode)

	pre.Next = rNode
	lNode.Next = succ
	return dummyHead.Next
}

func reverse(head *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
}

func main() {

}
