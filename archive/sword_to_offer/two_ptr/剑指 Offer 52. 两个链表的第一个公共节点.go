package two_ptr

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	markA, markB := false, false
	if headA == nil || headB == nil {
		return nil
	}
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
		if p1 == nil {
			if markA {
				return nil
			}
			p1 = headB
			markA = true
		}
		if p2 == nil {
			if markB {
				return nil
			}
			p2 = headA
			markB = true
		}
	}
	return p1
}
