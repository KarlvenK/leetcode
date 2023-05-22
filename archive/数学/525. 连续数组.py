from typing import *

class Solution:
    def findMaxLength(self, nums: List[int]) -> int:
        # prefix = num_1 - num_0
        prefix = 0
        hash = {
            0: -1,
        }
        ans = 0
        for i in range(len(nums)):
            if nums[i] == 0:
                prefix -= 1
            else:
                prefix += 1
            
            if prefix in hash:
                ans = max(ans, i - hash[prefix])
            else:
                if prefix not in hash:
                    hash[prefix] = i

        return ans    

def main():
    s = Solution()
    nums = [1,1,0,1,0,0,1]
    print(s.findMaxLength(nums))

if __name__ == "__main__":
    main()