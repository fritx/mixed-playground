# 1518. 换水问题
# https://leetcode.cn/problems/water-bottles/description/

# 方法二：数学，O(1)
class Solution:
    def numWaterBottles(self, numBottles: int, numExchange: int) -> int:
        b = numBottles
        ex, gain = numExchange, 1
        if b < ex: return b

        # 答案ans = b + n
        # 推导：首次打破 b - n(ex-gain) >= e 的n是多少
        # b - n(ex-gain) < ex
        # n > (b-ex) / (ex-gain)  # 接下来是关键
        # n_min = floor( 1 + ((b-ex) / (ex-gain)) )
        #       = 1 + (b-ex) // (ex-gain)
        n = floor( 1 + (b-ex) / (ex-gain) )
        return b + n

# 方法一：模拟，O(b-e)
class Solution:
    def numWaterBottles(self, numBottles: int, numExchange: int) -> int:
        bottle, ans = numBottles, numBottles
        while bottle >= numExchange:
            bottle -= numExchange
            ans += 1
            bottle += 1
        return ans
