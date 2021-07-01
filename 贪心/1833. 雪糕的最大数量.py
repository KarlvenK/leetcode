from typing import *

class Solution:
    def maxIceCream(self, costs: List[int], coins: int) -> int:
        costs.sort()
        ans = 0
        for price in costs:
            if coins >= price:
                ans += 1
                coins -= price
        return ans