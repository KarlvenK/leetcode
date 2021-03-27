package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	start := head
	n := 0
	tail := &ListNode{-1, nil}
	for head != nil {
		n++
		tail = head
		head = head.Next
	}
	k = k % n
	tail.Next = start
	ans := n - k
	pre := tail
	for i := 0; i < ans; i++ {
		pre = start
		start = start.Next
	}
	pre.Next = nil
	return start
}
