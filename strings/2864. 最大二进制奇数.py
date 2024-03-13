# 2864. 最大二进制奇数
# https://leetcode.cn/problems/maximum-odd-binary-number/

class Solution:
    def maximumOddBinaryNumber(self, s: str) -> str:

        # 题目OJ能通过 但本地测试用例不通过
        if s == "": return ""

        cnt = s.count('1')
        return '1' * (cnt - 1) + '0' * (len(s) - cnt) + '1'
