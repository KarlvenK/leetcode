package two_ptr

func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	slow := head
	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
