# 1969. 数组元素的最小非零乘积
# https://leetcode.cn/problems/minimum-non-zero-product-of-the-array-elements/

class Solution:
    def minNonZeroProduct(self, p: int) -> int:
        if p == 1:
            return 1
        # ** bug 3 - `1e9` not `10e9`
        # mod = 10e9+7
        # ** bug 4 - `1e9` is float not int
        # mod = 1e9+7
        # mod = int(1e9+7)
        mod = 10**9+7
        x = self.fastPow(2, p, mod)
        # ** bug 2 - should not mod at `y`
        # y = self.fastPow(2, p-1, mod)
        y = 1 << (p-1)
        # = (2^p-2)^(2^(p-1)-1) * (2^p-1)
        return self.fastPow(x-2, y-1, mod) * (x-1) % mod

    def fastPow(self, x: int, n: int, mod: int) -> int:
        res = 1
        while n > 0:
            # ** bug 1 - missing condition
            if n & 1 == 1:
                res = res * x % mod
            x = (x * x) % mod
            n >>= 1
        return res
