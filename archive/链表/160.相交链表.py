from typing import *

class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def getIntersectionNode(self, headA: ListNode, headB: ListNode) -> ListNode:
        if headA == None or headB == None:
            return None
        
        p1, p2 = headA, headB
        while p1 != p2:
            if p1 == None:
                p1 = headB
            else:
                p1 = p1.next
            if p2 == None:
                p2 = headA
            else:
                p2 = p2.next
        return p1
"""
  A:  0 -> 1 -> 2 -> 3
                      \
                        > 4 -> 5 -> 6
                      /
  B:       7 -> 8 -> 9

4-5-6是共同部分

0-1-2-3-4-5-6-7-8-9  - 4-5-6
7-8-9-4-5-6-0-1-2-3  - 4-5-6
若A B 有相交的地方，则A走完自己的链表后继续走B，B同理，同时进行。
最后他们走完的距离是一样。再走一步就会相遇。

"""