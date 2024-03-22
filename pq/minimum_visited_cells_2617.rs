use std::cmp::Reverse;
use std::collections::BinaryHeap;

pub struct Solution;

impl Solution {
    pub fn minimum_visited_cells(grid: Vec<Vec<i32>>) -> i32 {
        let m = grid.len();
        let n = grid[0].len();
        let mut dist = vec![vec![-1; n]; m];
        dist[0][0] = 1;
        let mut row: Vec<BinaryHeap<(Reverse<i32>, usize)>> = vec![BinaryHeap::new(); m];
        let mut col: Vec<BinaryHeap<(Reverse<i32>, usize)>> = vec![BinaryHeap::new(); n];

        let update = |x: i32, y: i32| {
            if x == -1 || y < x {
                return y;
            } else {
                return x;
            }
        };
        for i in 0..m {
            for j in 0..n {
                while let Some((_, x)) = row[i].peek() {
                    if *x as i32 + grid[i][*x] < j as i32 {
                        row[i].pop();
                    } else {
                        break;
                    }
                }
                if let Some((_, x)) = row[i].peek() {
                    dist[i][j] = update(dist[i][j], dist[i][*x] + 1);
                }
                while let Some((_, x)) = col[j].peek() {
                    if *x as i32 + grid[*x][j] < i as i32 {
                        col[j].pop();
                    } else {
                        break;
                    }
                }
                if let Some((_, x)) = col[j].peek() {
                    dist[i][j] = update(dist[i][j], dist[*x][j] + 1);
                }
                if dist[i][j] != -1 {
                    row[i].push((Reverse(dist[i][j]), j));
                    col[j].push((Reverse(dist[i][j]), i));
                }
            }
        }
        dist[m - 1][n - 1]
    }
}
