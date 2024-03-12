struct Solution;

impl Solution {
    pub fn capitalize_title(title: String) -> String {
        let words: Vec<&str> = title.split(' ').collect();
        let mut res = Vec::new();
        for word in words {
            if word.len() > 2 {
                res.push(word[..1].to_uppercase() + &word[1..].to_lowercase());
            } else {
                res.push(word.to_lowercase());
            }
        }
        res.join(" ")
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
            ["capiTalIze tHe titLe", "Capitalize The Title"],
            ["First leTTeR of EACH Word", "First Letter of Each Word"],
            ["i lOve leetcode", "i Love Leetcode"]
        ]"#;
    let cases: Vec<Case> = serde_json::from_str(json).unwrap();
    for case in cases {
        let ans = Solution::capitalize_title(case.input);
        assert_eq!(ans, case.want);
    }
}
