// 2369. 检查数组是否存在有效划分
// https://leetcode.cn/problems/check-if-there-is-a-valid-partition-for-the-array/

// 动态规划 / O(n), O(n)
struct Solution;
impl Solution {
    pub fn valid_partition(nums: Vec<i32>) -> bool {
        let n = nums.len();
        let mut dp = vec![false; n + 1];
        dp[0] = true;
        for i in 1..=n {
            if i >= 2 {
                dp[i] = dp[i - 2] && Self::valid_two(nums[i - 2], nums[i - 1]);
            }
            if i >= 3 {
                dp[i] =
                    dp[i] || dp[i - 3] && Self::valid_three(nums[i - 3], nums[i - 2], nums[i - 1]);
            }
        }
        dp[n]
    }
    fn valid_two(num1: i32, num2: i32) -> bool {
        num1 == num2
    }
    fn valid_three(num1: i32, num2: i32, num3: i32) -> bool {
        num1 == num2 && num1 == num3 || num1 == num2 - 1 && num1 == num3 - 2
    }
}

// 递归 / O(n), O(n)
struct Solution2;
impl Solution2 {
    pub fn valid_partition(nums: Vec<i32>) -> bool {
        let mut mem = vec![0; 100001];
        Self::dfs(0, &nums, &mut mem)
    }
    fn dfs(idx: usize, nums: &Vec<i32>, mem: &mut Vec<i32>) -> bool {
        let nums_size = nums.len();
        if nums_size - idx < 2 {
            return false;
        }
        if nums_size - idx <= 3 {
            return match nums_size - idx {
                2 => nums[idx] == nums[idx + 1],
                3 => {
                    (nums[idx] == nums[idx + 1] && nums[idx] == nums[idx + 2])
                        || (nums[idx] + 1 == nums[idx + 1] && nums[idx] + 2 == nums[idx + 2])
                }
                _ => false,
            };
        }
        if mem[idx] != 0 {
            return mem[idx] == 1;
        }
        let res = if nums[idx] == nums[idx + 1] {
            Self::dfs(idx + 2, nums, mem)
                || (nums[idx] == nums[idx + 2] && Self::dfs(idx + 3, nums, mem))
        } else {
            nums[idx] + 1 == nums[idx + 1]
                && nums[idx] + 2 == nums[idx + 2]
                && Self::dfs(idx + 3, nums, mem)
        };
        mem[idx] = if res { 1 } else { -1 };
        res
    }
}

#[test]
fn it_works() {
    assert_eq!(Solution::valid_partition(vec![4, 4, 4, 5, 6]), true);
    assert_eq!(Solution::valid_partition(vec![1, 1, 1, 2]), false);
    assert_eq!(Solution2::valid_partition(vec![4, 4, 4, 5, 6]), true);
    assert_eq!(Solution2::valid_partition(vec![1, 1, 1, 2]), false);
}
