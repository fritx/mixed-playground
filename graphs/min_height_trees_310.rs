use std::collections::VecDeque;
use std::vec::Vec;

pub struct Solution;

// 方法三：拓扑排序
impl Solution {
    // fn find_min_height_trees(n: i32, edges: Vec<(i32, i32)>) -> Vec<i32> {
    fn find_min_height_trees(n: i32, edges: Vec<Vec<i32>>) -> Vec<i32> {
        if n == 1 {
            return vec![0];
        }

        let mut graph = vec![Vec::new(); n as usize];
        let mut degrees = vec![0; n as usize];
        for edge in edges {
            // to access tuple elements, use: `.0`rustc(E0608)
            let x = edge[0] as usize;
            let y = edge[1] as usize;
            graph[x].push(y);
            graph[y].push(x);
            degrees[x] += 1;
            degrees[y] += 1;
        }

        let mut leaves = VecDeque::new();
        for (i, degree) in degrees.iter().enumerate() {
            if *degree == 1 {
                leaves.push_back(i as i32);
            }
        }

        let mut remaining_nodes = n;
        while remaining_nodes > 2 {
            let len = leaves.len();
            remaining_nodes -= len as i32;

            for _ in 0..len {
                let leaf = leaves.pop_front().unwrap();
                let neighbors = &graph[leaf as usize];
                // fix: type `usize` cannot be dereferenced - rustc
                // for &neighbor in neighbors {
                for neighbor in neighbors {
                    degrees[*neighbor] -= 1;
                    if degrees[*neighbor] == 1 {
                        leaves.push_back(*neighbor as i32);
                    }
                }
            }
        }
        leaves.into_iter().collect()
    }
}

#[allow(unused_imports)]
// use crate::utils::assert::assert_arrays_equal_unordered_2;
use crate::utils::assert::assert_arrays_equal_unordered;

use serde::Deserialize;

#[derive(Deserialize)]
struct Case {
    n: i32,
    // edges: Vec<(i32, i32)>,
    edges: Vec<Vec<i32>>,
    want: Vec<i32>,
}

#[test]
fn it_works() {
    let json = r#"[
            [4, [[1,0],[1,2],[1,3]], [1]],
            [6, [[3,0],[3,1],[3,2],[3,4],[5,4]], [3,4]],
            [2, [[0,1]], [0,1]],
            [1, [], [0]],
            [0, [], []]
        ]"#;
    let cases: Vec<Case> = serde_json::from_str(json).unwrap();

    for case in cases {
        let ans = Solution::find_min_height_trees(case.n, case.edges);
        // assert_eq!(ans, case.want);
        assert_arrays_equal_unordered(&ans, &case.want);
    }
}
