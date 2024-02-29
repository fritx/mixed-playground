// 2673. 使二叉树所有路径值相等的最小代价
// https://leetcode.cn/problems/make-costs-of-paths-equal-in-a-binary-tree/

#[cfg(test)]
struct Solution;

#[cfg(test)]
impl Solution {
    pub fn min_increments(n: i32, cost: Vec<i32>) -> i32 {
        let mut ans = 0;
        let mut cost = cost.clone();
        for i in (0..(n as usize) - 1).rev().step_by(2) {
            ans += (cost[i] - cost[i + 1]).abs();
            // 叶节点 i 和 i+1 的双亲节点下标为 i/2（整数除法）
            cost[i / 2] += cost[i].max(cost[i + 1]);
        }
        ans
    }
}

#[test]
fn it_works() {
    assert_eq!(Solution::min_increments(7, vec![1, 5, 2, 2, 3, 3, 1]), 6);
    assert_eq!(Solution::min_increments(3, vec![5, 3, 3]), 0);
}
