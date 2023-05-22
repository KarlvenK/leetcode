package interview

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if slow == fast {
			p := head
			for p != fast {
				p = p.Next
				fast = fast.Next
			}
			return p
		}
	}
	return nil
}

/*
设一快一慢指针， 有环（顺时针）链表
		            /----b-----\
（head）----a-------	            x
		            \------c--/
fast 和 slow相遇在环中x处时，x将环分成b、c两部分
由于fast经过路程为slow 的两倍
那么可以得到下列等式：
a + (n + 1)b + nc  = 2(a + b)    (n >= 1)
化简得：
(n - 1)(b + c) + c = a
从而，我们另起一个新节点p = head，让它和x处的点同步遍历下去，当他们相遇时，刚好就是环的起点。
 */
