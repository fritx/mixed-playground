use serde::Deserialize;

// fix: unused import: `crate::graphs::count_paths_1976::Solution`
// `#[warn(unused_imports)]` on by default
#[allow(unused_imports)]
// use crate::graphs::count_paths_1976::Solution;
use super::count_paths_1976::Solution;

// 定义结构体来表示最外层的元素
#[derive(Deserialize)]
struct Case {
    n: i32,
    // roads: Vec<(i32, i32, i32)>,
    roads: Vec<Vec<i32>>,
    want: i32,
}

#[test]
fn it_works() {
    // JSON字符串
    let json = r#"[
            [7, [[0,6,7],[0,1,2],[1,2,3],[1,3,3],[6,3,3],[3,5,1],[6,5,1],[2,5,1],[0,4,5],[4,6,2]], 4],
            [2, [[1,0,10]], 1],
            [2, [], 0]
        ]"#;

    // 使用serde_json解析JSON字符串
    // let result: Result<Vec<Case>, serde_json::Error> = serde_json::from_str(json);
    // let cases = result.unwrap();
    let cases: Vec<Case> = serde_json::from_str(json).unwrap();

    // 如果解析成功，处理解析后的数据
    for case in cases {
        let ans = Solution::count_paths(case.n, case.roads);
        assert_eq!(ans, case.want);
    }
}
