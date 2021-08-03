package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scanln(&n)
	head := &ListNode{
		-1,
		nil,
	}
	pre := head
	for i := 0; i < n; i++ {
		var temp int
		_, _ = fmt.Scanln(&temp)
		cur := &ListNode{
			temp,
			nil,
		}
		pre.Next = cur
		pre = cur
	}
	auto := deleteDuplicates(head.Next)
	for auto != nil {
		fmt.Println(auto.Val)
		auto = auto.Next
	}
}

func _deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{
		-100000,
		nil,
	}
	pre := dummyHead
	cur := head

	for cur != nil && cur.Next != nil { //可能产生只有cur没有cur.Next的情况
		if cur.Val == cur.Next.Val {
			mark := cur.Val
			for cur != nil && cur.Val == mark {
				cur = cur.Next
			}
		} else {
			pre.Next = cur
			pre = cur
			cur = cur.Next
			pre.Next = nil // 切断 ！！！！！
		}
	}
	if cur != nil { //最后还剩单个
		pre.Next = cur
	}
	return dummyHead.Next
}
