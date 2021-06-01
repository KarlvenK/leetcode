from typing import List

class Solution:
    def canEat(self, candiesCount: List[int], queries: List[List[int]]) -> List[bool]:
        prefix_sum = list()
        prefix_sum.append(0)
        for i in range(len(candiesCount)):
            prefix_sum.append(prefix_sum[-1] + candiesCount[i])

        ans = list()
        for fType, fDay, dCap in queries:
            a1 = fDay + 1
            a2 = dCap * a1

            b1 = 1 if fType == 0 else prefix_sum[fType] + 1
            b2 = prefix_sum[fType + 1]

            if a1 >b2 or a2 < b1:
                ans.append(False)
            else:
                ans.append(True)
        
        return ans
            
            



"""
    第i天可能吃到的糖果的编号范围：
    [ favorateDayi + 1,  dailyCapi * (favorateDayi + 1) ]
    用前缀和算出每种糖果对应的编号区间
    [ ]
    对比一下区间是否重合就可以知道答案
"""
