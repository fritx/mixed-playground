// 1883. 准时抵达会议现场的最小跳过休息次数
// https://leetcode.cn/problems/minimum-skips-to-arrive-at-meeting-on-time/

use std::f64::INFINITY;

// 方法一：动态规划
// 注意 IEEE-754标准浮点数精度问题 可忽略误差处理
impl Solution {
    pub fn min_skips(dist: Vec<i32>, speed: i32, hours_before: i32) -> i32 {
        // 可忽略误差
        let eps = 1e-7;
        let n = dist.len();
        let mut f = vec![vec![f64::INFINITY; n + 1]; n + 1];
        f[0][0] = 0.0;
        for i in 1..= n {
            for j in 0..= i {
                if j != i {
                    f[i][j] = f[i][j].min((f[i - 1][j] + (dist[i - 1] as f64) / (speed as f64) - eps).ceil());
                }
                if j != 0 {
                    f[i][j] = f[i][j].min(f[i - 1][j - 1] + (dist[i - 1] as f64) / (speed as f64));
                }
            }
        }
        for j in 0..= n {
            if f[n][j] < (hours_before as f64) + eps {
                return j as i32;
            }
        }
        -1
    }
}
