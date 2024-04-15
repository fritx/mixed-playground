//go:build ignore

// 1483. 树节点的第 K 个祖先
// https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/

// 1 <= k <= n <= 5 * 10^4
// parent[0] == -1 表示编号为 0 的节点是根节点。
// 对于所有的 0 < i < n ，0 <= parent[i] < n 总成立
// 0 <= node < n
// 至多查询 5 * 10^4 次

// 1. 模拟 时间O(n) 空间O(1) 8/17超出时间限制
// 2. 哈希表 时间O(1) 空间O(n^2) 7/17超出内存限制

// 3. 倍增
// 初始化 时间O(nlogn) 空间O(nlogn)
// 单次查询 时间O(logn) 空间O(1)
class TreeAncestor {
public:
    constexpr static int Log = 16;
    vector<vector<int>> ancestors;

    TreeAncestor(int n, vector<int>& parent) {
        ancestors = vector<vector<int>>(n, vector<int>(Log, -1));
        for (int i = 0; i < n; i++) {
            ancestors[i][0] = parent[i];
        }
        for (int j = 1; j < Log; j++) {
            for (int i = 0; i < n; i++) {
                if (ancestors[i][j - 1] != -1) {
                    ancestors[i][j] = ancestors[ancestors[i][j - 1]][j - 1];
                }
            }
        }
    }
    int getKthAncestor(int node, int k) {
        for (int j = 0; j < Log; j++) {
            if ((k >> j) & 1) {
                node = ancestors[node][j];
                if (node == -1) {
                    return -1;
                }
            }
        }
        return node;
    }
};
