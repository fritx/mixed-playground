# 2917. 找出数组中的 K-or 值
# https://leetcode.cn/problems/find-the-k-or-of-an-array/

from typing import List

# 方法一：枚举每一个二进制位
class Solution:
    def findKOr(self, nums: List[int], k: int) -> int:
        ans = 0
        for i in range(31):  # int32
            cnt = sum(1 for n in nums if (n >> i) & 1)
            if cnt >= k:
                ans |= 1 << i
        return ans

# 方法一：枚举每一个二进制位
# @author 文心一言
class Solution2:
    def findKOr(self, nums, k):
        # 确定数组 nums 的最大位数
        max_bit = max(nums).bit_length() - 1

        # 初始化 K-or 值为 0
        kor_value = 0

        # 遍历每一位
        for bit in range(max_bit + 1):
            # 初始化计数器为 0
            count = 0

            # 遍历数组 nums 中的每个元素
            for num in nums:
                # 检查当前元素的当前位是否为 1
                if (num >> bit) & 1:  # (num >> bit) 将 num 右移 bit 位，然后 & 1 检查最低位是否为 1
                    count += 1

            # 如果当前位上至少有 k 个 1，则设置 K-or 值的该位为 1
            if count >= k:
                kor_value |= (1 << bit)  # (1 << bit) 创建一个只有第 bit 位为 1 的数，然后与 kor_value 进行按位或运算

        return kor_value

# 路人
class Solution3:
    def findKOr(self, nums, k):
        # 使用一个Python字典来记录每个位上1的数量
        bit_counts = {i: 0 for i in range(32)}

        # 遍历nums中的每个数字
        for num in nums:
            # 遍历数字的每一位
            for i in range(32):
                # 检查当前位是否为1
                if num & (1 << i):
                    bit_counts[i] += 1

        # 构建K-or值
        kor_value = 0
        for i in range(32):
            # 如果当前位上1的数量大于等于k，则设置K-or值的相应位为1
            if bit_counts[i] >= k:
                kor_value |= (1 << i)

        return kor_value
