struct Solution;

impl Solution {
    pub fn maximum_odd_binary_number(s: String) -> String {
        // fix: attempt to subtract with overflow
        // 因为后面用到了 (usize) cnt-1 所以溢出了
        // 题目OJ能通过 但本地测试用例不通过
        if s.len() == 0 {
            return "".to_string();
        }
        let cnt = s.chars().filter(|&c| c == '1').count();
        "1".repeat(cnt - 1) + &*"0".repeat(s.len() - cnt) + "1"
    }
}

use serde::Deserialize;

// 定义结构体来表示最外层的元素
#[derive(Deserialize)]
struct Case {
    input: String,
    want: String,
}

#[test]
fn it_works() {
    let json = r#"[
            ["010", "001"],
            ["0101", "1001"],
            ["1", "1"],
            ["", ""]
        ]"#;
    let cases: Vec<Case> = serde_json::from_str(json).unwrap();
    for case in cases {
        let ans = Solution::maximum_odd_binary_number(case.input);
        assert_eq!(ans, case.want);
    }
}
