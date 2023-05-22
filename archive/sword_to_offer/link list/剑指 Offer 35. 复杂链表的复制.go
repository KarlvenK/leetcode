package link_list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	dict := make(map[*Node]int)
	fakeHead := head
	for i := 0; fakeHead != nil; i++ {
		dict[fakeHead] = i
		fakeHead = fakeHead.Next
	}
	dummy := &Node{}
	//fakeHead = head
	newDict := make(map[int]*Node)
	var cur *Node
	pre := dummy
	fakeHead = head

	for i := 0; head != nil; i++ {
		cur = &Node{
			Val: head.Val,
		}
		newDict[i] = cur
		pre.Next = cur
		pre = cur
		head = head.Next
	}
	head = fakeHead
	nhead := dummy.Next
	for head != nil {
		if head.Random != nil {
			nhead.Random = newDict[dict[head.Random]]
		}
		nhead = nhead.Next
		head = head.Next
	}
	return dummy.Next
}
