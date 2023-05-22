from typing import *

class Solution:
    def checkSubarraySum(self, nums: List[int], k: int) -> bool:
        """
        hash_table = {
            nums[0]: 0,
        }
        prefix_sum = [nums[0]]
        for i in range(1,len(nums)):
            prefix_sum.append(prefix_sum[i - 1] + nums[i])
            hash_table[prefix_sum[i]] = i
        
        for i in range(1, len(nums)):
            sum = self.getLim(prefix_sum[i], k) 
            if sum == prefix_sum[i]:
                return True   
            # print(sum)
            if sum < k:
                break
            while True:
                v = prefix_sum[i] - sum
                if v in hash_table:
                    return True
                sum -= k
                if sum < nums[i]:
                    break
        
        return False
        
        def getLim(self, val, k : int) -> int:
            return (val // k) * k        
        """
        n = len(nums)
        if n < 2:
            return False

        map = {
            0: -1,
        }
        sum = 0
        for i in range(0, n):
            sum += nums[i]
            sum = sum % k
            if sum in map:
                if i - map[sum] > 1:
                    return True
            else:
                map[sum] = i

        return False



def main():
    nums = [23,2,4,6,6]
    s = Solution()
    k = 7
    print(s.checkSubarraySum(nums, k))

    nums = [2, 1]
    k = 2
    print(s.checkSubarraySum(nums, k))
    nums = [5, 0, 0, 0]
    k = 3
    print(s.checkSubarraySum(nums, k))

if __name__ == "__main__":
    main()

"""
同余定理：如果两个整数m、n满足n-m能被k整除，那么就成n和m对k同余
prefix[x] = sum[0..x]

n = len(nums)
０ <= i <　ｊ < n
prefix[j] - prefix[i] = k * n
只要prefix[i] 与 prefix[j] 同余就行
所以可以用哈希表保存 某一种对k余数 出现的最早位置就行

进一步，prefix只要存前缀和的余数就行
再进一步，prefix数组可以被迭代掉

最后当出现同余时，判断j - i是否大于等于2就行

"""