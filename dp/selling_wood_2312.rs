// 2312. 卖木头块
// https://leetcode.cn/problems/selling-pieces-of-wood/

use std::collections::HashMap;

struct Solution;

// 方法一：动态规划 / 记忆化搜索
impl Solution {
    pub fn selling_wood(m: i32, n: i32, prices: Vec<Vec<i32>>) -> i64 {
        let mut value: HashMap<(i32, i32), i32> = HashMap::new();
        let mut memo: Vec<Vec<i64>> = vec![vec![-1; (n + 1) as usize]; (m + 1) as usize];

        fn dfs(x: i32, y: i32, memo: &mut Vec<Vec<i64>>, value: &HashMap<(i32, i32), i32>) -> i64 {
            if memo[x as usize][y as usize] != -1 {
                return memo[x as usize][y as usize];
            }
            let mut ret = if value.contains_key(&(x, y)) {
                value[&(x, y)] as i64
            } else {
                0
            };

            if x > 1 {
                for i in 1..x {
                    ret = ret.max(dfs(i, y, memo, value) + dfs(x - i, y, memo, value));
                }
            }
            if y > 1 {
                for j in 1..y {
                    ret = ret.max(dfs(x, j, memo, value) + dfs(x, y - j, memo, value));
                }
            }
            memo[x as usize][y as usize] = ret;
            ret
        }

        for price in prices {
            value.insert((price[0], price[1]), price[2]);
        }
        dfs(m, n, &mut memo, &value)
    }
}

use serde::Deserialize;

// 定义结构体来表示最外层的元素
#[derive(Deserialize)]
struct Case {
    m: i32,
    n: i32,
    prices: Vec<Vec<i32>>,
    want: i64,
}

#[test]
fn it_works() {
    let json = r#"[
            [3, 5, [[1,4,2],[2,2,7],[2,1,3]], 19],
            [4, 6, [[3,2,10],[1,4,2],[4,1,3]], 32],
            [7, 7, [], 0],
            [0, 0, [[1,2,3]], 0],
            [0, 0, [], 0]
        ]"#;
    let cases: Vec<Case> = serde_json::from_str(json).unwrap();

    for case in cases {
        let ans = Solution::selling_wood(case.m, case.n, case.prices);
        assert_eq!(ans, case.want);
    }
}
