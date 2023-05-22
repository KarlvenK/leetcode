from typing import *

class Solution:
    def frequencySort(self, s: str) -> str:
        cnt = {}
        for i in range(len(s)):
            if cnt.__contains__(s[i]):
                cnt[s[i]] += 1
            else:
                cnt.setdefault(s[i], 1)
        lim = 0
        tool = [[] for i in range(len(s) + 1)]
        for key, value in cnt.items():
            tool[value].append(key)
            lim = max(lim, value)
        
        ans = ''
        for i in range(lim, 0, -1):
            for j in range(len(tool[i])):
                for k in range(i):
                    ans += tool[i][j]
        return ans
        