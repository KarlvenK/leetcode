class Solution:
    def isPowerOfFour(self, n: int) -> bool:
        return n > 0 and (n & (n - 1)) == 0 and (n & 0xaaaaaaaa) == 0

"""
给定一个整数，写一个函数来判断它是否是 4 的幂次方。如果是，返回 true ；否则，返回 false 。
整数 n 是 4 的幂次方需满足：存在整数 x 使得 n == 4x

solution:
4的幂次方是2的幂次方， 且必是 "1（若干个0）0" 的形式

10000000 // n
01111111 // n - 1
相与必为0

4的幂的1的位置必然在奇数位
1010 1010 1010 1010 1010 1010 1010 1010
  a               ...               a

"""