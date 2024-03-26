# 2642. 设计可以求最短路径的图类
# https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/

import heapq
from math import inf
from typing import List

# 方法一：Dijkstra 求最短路径
# 其中 m 表示给定的 edges 数组的长度，k 表示调用 addEdge 的次数
# SC: O(n+m+k)
class Graph:
    # O(m)
    def __init__(self, n: int, edges: List[List[int]]):
        self.graph = [[] for _ in range(n)]
        for x, y, cost in edges:
            self.graph[x].append((y, cost))

    # O(1)
    def addEdge(self, edge: List[int]) -> None:
        x, y, cost = edge[0], edge[1], edge[2]
        self.graph[x].append((y, cost))

    # O((m+k)log(m+k))
    def shortestPath(self, node1: int, node2: int) -> int:
        dist = [inf] * len(self.graph)
        dist[node1] = 0
        q = [(0, node1)]
        while q:
            cost, x = heapq.heappop(q)
            if x == node2:
                return cost
            for y, ncost in self.graph[x]:
                if dist[y] > cost + ncost:
                    dist[y] = cost + ncost
                    heapq.heappush(q, [cost + ncost, y])
        return -1

# 方法二：Floyd 求最短路径
# SC: O(n^2)
class Graph_2:
    # O(n^3+m)
    def __init__(self, n: int, edges: List[List[int]]):
        self.dist = [[inf] * n for _ in range(n)]
        for i in range(n):
            self.dist[i][i] = 0
        for x, y, cost in edges:
            self.dist[x][y] = cost
        for k in range(n):
            for i in range(n):
                for j in range(n):
                    self.dist[i][j] = min(self.dist[i][j], self.dist[i][k] + self.dist[k][j])

    # O(n^2)
    def addEdge(self, edge: List[int]) -> None:
        x, y, cost = edge[0], edge[1], edge[2]
        if cost >= self.dist[x][y]:
            return
        n = len(self.dist)
        for i in range(n):
            for j in range(n):
                self.dist[i][j] = min(self.dist[i][j], self.dist[i][x] + cost + self.dist[y][j])

    # O(1)
    def shortestPath(self, node1: int, node2: int) -> int:
        res = self.dist[node1][node2]
        return res if res != inf else -1

# 自己的Floyd
class Graph_3:
    def __init__(self, n: int, edges: List[List[int]]):
        self.dist = [[inf] * n for _ in range(n)]
        # ** bug 1.1 - missing assigning `self.dist` to `dist`
        dist = self.dist
        # ** bug 2.1 - missing defining `n`
        self.n = n
        for i in range(n):
            dist[i][i] = 0
        for x, y, cost in edges:
            dist[x][y] = min(dist[x][y], cost)
        # ** bug 4 - for `k,i,j` instead of `i,j,k` ????
        # for i in range(n):
        #     for j in range(n):
        #         for k in range(n):
        for k in range(n):
            for i in range(n):
                for j in range(n):
                    dist[i][j] = min(dist[i][j], dist[i][k] + dist[k][j])

    def addEdge(self, edge: List[int]) -> None:
        # ** bug 1.2 - missing assigning `self.dist` to `dist`
        dist = self.dist
        x, y, cost = edge
        if cost >= dist[x][y]: return
        # ** bug 2.2 - missing defining `n`
        n = self.n
        for i in range(n):
            for j in range(n):
                dist[i][j] = min(dist[i][j], dist[i][x] + cost + dist[y][j])

    def shortestPath(self, node1: int, node2: int) -> int:
        # ** bug 3 - missing returning `-1` if `inf`
        # return self.dist[node1][node2]
        res = self.dist[node1][node2]
        return -1 if res == inf else res
