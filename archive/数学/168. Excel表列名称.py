from collections import defaultdict
from typing import *

class Solution:
    def convertToTitle(self, columnNumber: int) -> str:
        ans = ''
        while columnNumber > 0:
            mod = (columnNumber - 1) % 26 + 1          #*
            ans = chr(ord('A') + mod - 1) + ans
            columnNumber = (columnNumber - mod) // 26  #**
        return ans

def main():
    s = Solution()
    for i in range(1, 26 * 3):
        print(s.convertToTitle(i))

main()
'''
                  n
columnNumber = sigama (26^i * ai) = "a(n) - a(n-1) - a(n - 2) ... a(0)"
                i = 0
且 1 <= ai <= 26
每次对26取余数来算出ai，但是由于ai的范围是[1,26]所以进行微调，当ai是26时，应当让它不直接归零（#* 行）
最后减去这个ai再除以26来计算ai+1
由于ai范围[1,26]，不减去ai直接除会导致结果多 1 （26/26 = 1） （#** 行）
'''