package main

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = prev;
		prev = cur
		cur = temp
	}
	return prev
}
func main() {
	
}
