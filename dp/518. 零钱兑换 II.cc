//go:build ignore

// 518. 零钱兑换 II
// https://leetcode.cn/problems/coin-change-ii/

// 方法一：动态规划
class Solution {
public:
    int change(int amount, vector<int>& coins) {
        vector<int> dp(amount + 1);
        dp[0] = 1;
        for (int& coin : coins) {
            for (int i = coin; i <= amount; i++) {
                dp[i] += dp[i - coin];
            }
        }
        return dp[amount];
    }
};
