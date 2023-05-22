from typing import *
class Solution:
    def findErrorNums(self, nums: List[int]) -> List[int]:
        dic = {}
        ans = []
        for num in nums:
            if num not in dic:
                dic.setdefault(num, 1)
            else:
                ans.append(num)
        
        for i in range(1, len(nums)+1):
            if i not in dic:
                ans.append(i)
                break
        return ans
