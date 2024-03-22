#[allow(unused_imports)]
use super::minimum_visited_cells_2617::Solution;

use serde::Deserialize;

// 定义结构体来表示最外层的元素
#[derive(Deserialize)]
struct Case {
    grid: Vec<Vec<i32>>,
    want: i32,
}

#[test]
fn it_works() {
    let json_string = r#"
        [
            [[[3,4,2,1],[4,2,3,1],[2,1,0,0],[2,4,0,0]], 4],
            [[[3,4,2,1],[4,2,1,1],[2,1,1,0],[3,4,1,0]], 3],
            [[[2,1,0],[1,0,0]], -1]
        ]"#;
    let cases: Vec<Case> = serde_json::from_str(json_string).unwrap();

    for case in cases.iter() {
        // fix: cannot move out of `case.grid` which is behind a shared reference
        //      move occurs because `case.grid` has type `Vec<Vec<i32>>`, which does not implement the `Copy` trait
        // let ans = Solution::minimum_visited_cells(case.grid);
        let ans = Solution::minimum_visited_cells(case.grid.clone());
        assert_eq!(ans, case.want);
    }
}
