use std::cmp::Reverse;
use std::collections::BinaryHeap;

// 方法一：Dijkstra 求最短路径
struct Graph {
    graph: Vec<Vec<(i32, i32)>>,
}
impl Graph {
    fn new(n: i32, edges: Vec<Vec<i32>>) -> Self {
        let mut graph = vec![Vec::new(); n as usize];
        for edge in edges {
            let x = edge[0];
            let y = edge[1];
            let cost = edge[2];
            graph[x as usize].push((y, cost));
        }
        Graph { graph }
    }
    fn add_edge(&mut self, edge: Vec<i32>) {
        let x = edge[0];
        let y = edge[1];
        let cost = edge[2];
        &self.graph[x as usize].push((y, cost));
    }
    fn shortest_path(&self, node1: i32, node2: i32) -> i32 {
        let mut pq = BinaryHeap::new();
        let mut dist = vec![std::i32::MAX; self.graph.len()];
        dist[node1 as usize] = 0;
        pq.push((Reverse(0), node1));
        while let Some((Reverse(cost), cur)) = pq.pop() {
            if cur == node2 {
                return cost;
            }
            for &(next, ncost) in &self.graph[cur as usize] {
                if dist[next as usize] > cost + ncost as i32 {
                    dist[next as usize] = cost + ncost as i32;
                    pq.push((Reverse(cost + ncost as i32), next));
                }
            }
        }
        -1
    }
}

// 方法二：Floyd 求最短路径
use std::cmp::min;

struct Graph_2 {
    dist: Vec<Vec<i32>>,
}
impl Graph_2 {
    fn new(n: i32, edges: Vec<Vec<i32>>) -> Self {
        let n = n as usize;
        let mut dist = vec![vec![i32::MAX; n]; n];
        for i in 0..n {
            dist[i][i] = 0;
        }
        for edge in edges {
            dist[edge[0] as usize][edge[1] as usize] = edge[2];
        }
        for k in 0..n {
            for i in 0..n {
                for j in 0..n {
                    if dist[i][k] != i32::MAX && dist[k][j] != i32::MAX {
                        dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j]);
                    }
                }
            }
        }
        return Graph { dist };
    }
    fn add_edge(&mut self, edge: Vec<i32>) {
        let x = edge[0] as usize;
        let y = edge[1] as usize;
        let cost = edge[2];
        if cost >= self.dist[x][y] {
            return;
        }
        let n = self.dist.len();
        for i in 0..n {
            for j in 0..n {
                if self.dist[i][x] != i32::MAX && self.dist[y][j] != i32::MAX {
                    self.dist[i][j] =
                        min(self.dist[i][j], self.dist[i][x] + cost + self.dist[y][j]);
                }
            }
        }
    }
    fn shortest_path(&self, node1: i32, node2: i32) -> i32 {
        let res = self.dist[node1 as usize][node2 as usize];
        if res == i32::MAX {
            -1
        } else {
            res
        }
    }
}
