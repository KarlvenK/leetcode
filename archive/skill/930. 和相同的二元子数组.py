from typing import *

class Solution:
    def numSubarraysWithSum(self, nums: List[int], goal: int) -> int:
        tool = {}
        sum, ans = 0, 0
        for num in nums:
            if sum not in tool:
                tool.setdefault(sum, 1)
            else:
                tool[sum] += 1
            sum += num
            if (sum - goal) in tool:
                ans += tool[sum - goal]
        return ans

def main():
    lst = [1,0,1,0,1]
    goal = 2
    s = Solution()
    print(s.numSubarraysWithSum(lst, goal))

main()
