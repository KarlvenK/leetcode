from typing import *

class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        ans, cnt = -1, 0
        n = len(nums)
        for i in range(n):
            if cnt == 0:
                ans = nums[i]
                cnt = 1
            else:
                if ans == nums[i]:
                    cnt += 1
                else:
                    cnt -= 1
        
        cnt = 0
        for i in range(n):
            if nums[i] == ans:
                cnt += 1
        if (cnt * 2) <= n:
            ans = -1
        return ans


def main():
    nums = [1,2,2,3,3,3,3]
    s = Solution()
    print(s.majorityElement(nums))

if __name__ == '__main__':
    main()        

'''
博耶-摩尔多数投票算法
是一种用来寻找一组元素中占多数元素的常数空间级时间复杂度算法

过程：

    初始化元素m并给计数器i赋初值i = 0
    对于输入队列中每一个元素x：
        若i = 0, 那么 m = x and i = 1
        否则
            若m = x, 那么 i = i + 1
            否则 i = i − 1
    返回 m


即便输入序列没有多数元素，这一算法也会返回一个序列元素。
可通过第二次遍历来确定是否为多数元素

一个比较感性的认识：由于过半，该元素的出现次数count，
一定大于其他元素的出现次数之和。所以在count+1 与count-1
不断交替过程中，他能保留下来
'''