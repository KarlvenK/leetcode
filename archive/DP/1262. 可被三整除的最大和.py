from typing import *
class Solution:
    def maxSumDivThree(self, nums: List[int]) -> int:
        n = len(nums)
        f = [[0 for _ in range(3)] for _ in range(n + 1)]
        inf = 1234567890
        f[0][1] = -inf
        f[0][2] = -inf

        for i in range(n):
            if nums[i] % 3 == 0:
                f[i + 1][0] = f[i][0] + nums[i]
                f[i + 1][1] = f[i][1] + nums[i]
                f[i + 1][2] = f[i][2] + nums[i]
            elif nums[i] % 3 == 1:
                f[i + 1][0] = max(f[i][0], f[i][2] + nums[i])
                f[i + 1][1] = max(f[i][1], f[i][0] + nums[i])
                f[i + 1][2] = max(f[i][2], f[i][1] + nums[i])
            else:
                f[i + 1][0] = max(f[i][0], f[i][1] + nums[i])
                f[i + 1][1] = max(f[i][1], f[i][2] + nums[i])
                f[i + 1][2] = max(f[i][2], f[i][0] + nums[i])
        return f[n][0]
