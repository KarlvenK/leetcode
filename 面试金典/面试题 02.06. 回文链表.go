package interview

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	if head.Next == nil {
		return true
	}
	firstEnd := GetFirstEnd(head)
	secondStart := reverseList(firstEnd.Next)
	p, q := head, secondStart
	ret := true
	for q != nil {
		if p.Val != q.Val {
			ret = false
			break
		}
		p, q = p.Next, q.Next
	}
	reverseList(secondStart)
	return ret
}
func GetFirstEnd(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	pre, cur := nil,  head

	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
