from typing import *

class Solution:
    def numWays(self, n: int, relation: List[List[int]], k: int) -> int:
        dp = [[0 for j in range(n)] for i in range(k + 1)]
        dp[0][0] = 1
        for i in range(1, k + 1):
            for t in relation[:]:
                src, dst = t[0], t[1]
                dp[i][dst] += dp[i - 1][src]
        return dp[k][n - 1]

def main():
    n = 5
    relation = [[0,2],[2,1],[3,4],[2,3],[1,4],[2,0],[0,4]]
    k = 3
    s = Solution()
    print(s.numWays(n, relation, k))


main()
    