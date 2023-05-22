package main

/*
给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
k是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{ //哑巴节点，后趋永远指向链表头，用于返回答案
		Next: head,
	}
	pre := dummy
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		nex := tail.Next
		head, tail = reversePart(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return dummy.Next
}

func reversePart(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := (*ListNode)(nil)
	tail.Next = nil
	/*
		这里我们把取出来的长度为k的链表的头的前驱和尾的后趋都截断
		让这端链表彻底断开
		在反转完之后返回他新的头和尾
		然后，主函数再将之前保存的前驱和后趋再接上
	*/
	//下面就是普通的反转链表
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre, head
}
