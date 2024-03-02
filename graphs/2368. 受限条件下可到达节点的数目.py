from collections import defaultdict
from typing import List

# 并查集 / O(n.α(n)),O(n)
class Solution3:
    def reachableNodes(self, n: int, edges: List[List[int]], restricted: List[int]) -> int:
        # 适配增强测试用例`0, [], [], 0`
        if n == 0: return 0
        is_restricted = [0] * n
        for x in restricted:
            is_restricted[x] = 1

        uf = UnionFind(n)
        for v in edges:
            if is_restricted[v[0]] or is_restricted[v[1]]:
                continue
            uf.merge(v[0], v[1])
        return uf.count()

class UnionFind:
    def __init__(self, n):
        self.f = list(range(n))
        self.rank = [0] * n

    def merge(self, x, y):
        rx = self.find(x)
        ry = self.find(y)
        if rx != ry:
            if self.rank[rx] > self.rank[ry]:
                self.f[ry] = rx
            elif self.rank[rx] < self.rank[ry]:
                self.f[rx] = ry
            else:
                self.f[ry] = rx
                self.rank[rx] += 1

    def find(self, x):
        if x != self.f[x]:
            self.f[x] = self.find(self.f[x])
        return self.f[x]

    def count(self):
        cnt = 0
        rt = self.find(0)
        for i in range(len(self.f)):
            if rt == self.find(i):
                cnt += 1
        return cnt

# DFS / O(n),O(n)
class Solution2:
    def reachableNodes(self, n: int, edges: List[List[int]], restricted: List[int]) -> int:
        # ** 答案这里错了
        # is_restricted = [1] * n
        is_restricted = [0] * n
        for x in restricted:
            is_restricted[x] = 1
        g = [[] for _ in range(n)]
        for v in edges:
            g[v[0]].append(v[1])
            g[v[1]].append(v[0])

        cnt = 0
        def dfs(x, f):
            nonlocal cnt
            cnt += 1
            for y in g[x]:
                if y != f and not is_restricted[y]:
                    dfs(y, x)

        # 适配增强测试用例`0, [], [], 0`
        if n == 0: return 0
        dfs(0, -1)
        return cnt

# BFS 暴力 / O(n^3),O(n^2)
class Solution:
    def reachableNodes(self, n: int, edges: List[List[int]], restricted: List[int]) -> int:
        # 适配增强测试用例`0, [], [], 0`
        if n == 0: return 0
        # ** impr 1.1
        # 题目OJ能通过 但此处适配增强测试用例`1, [], [0], 0`
        # if n == 1: return 1
        bs = set(restricted)
        # ** impr 1.2
        if n == 1: return 0 if 0 in bs else 1

        nb = defaultdict(list)
        for x, y in edges:
            if x in bs or y in bs: continue
            nb[x].append(y)
            nb[y].append(x)

        rs = set([0])
        vis = set()
        q = [0]
        while q:
            nq = []
            for x in q:
                if x in vis: continue
                vis.add(x)
                for y in nb[x]:
                    nq.append(y)
                    rs.add(y)
            q = nq
        return len(rs)
