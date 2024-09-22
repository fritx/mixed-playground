# 2739. 总行驶距离
# https://leetcode.cn/problems/total-distance-traveled/

# 方法二：数学，O(1)
# 公式推导过程，参考 1518. 换水问题
class Solution:
    def distanceTraveled(self, mainTank: int, additionalTank: int) -> int:
        m, a = mainTank, additionalTank
        cost, gain = 5, 1
        if m < cost or a < gain: return 10 * m

        # ans = 10 * ( m + min( a, n * gain ) )
        # 求 n，使得首次打破 m - n * (cost - gain) < cost
        # => n > (m - cost) / (cost - gain)  # 接下来是关键
        # 即 n_min = floor( 1 + (m - cost) / (cost - gain) )
        #               => 1 + (m - cost) // (cost - gain)
        n = floor( 1 + (m - cost) / (cost - gain) )
        return 10 * ( m + min( a, n * gain ) )

# 方法一：模拟，O(n)或O(n//4)
class Solution:
    def distanceTraveled(self, mainTank: int, additionalTank: int) -> int:
        dis = 0
        cost = 0
        while mainTank > 0:
            dis += 10
            cost += 1
            mainTank -= 1
            if cost % 5 == 0 and additionalTank > 0:
                mainTank += 1
                additionalTank -= 1
        return dis

class Solution:
    def distanceTraveled(self, mainTank: int, additionalTank: int) -> int:
        dis = 0
        while mainTank >= 5:
            dis += 50
            mainTank -= 5
            if additionalTank > 0:
                mainTank += 1
                additionalTank -= 1
        dis += 10 * mainTank
        return dis
