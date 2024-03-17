// 310. 最小高度树
// https://leetcode.cn/problems/minimum-height-trees/
// 1 <= n <= 2 * 10^4
// edges.length == n - 1
// 0 <= ai, bi < n
// ai != bi
// 所有 (ai, bi) 互不相同
// 给定的输入 保证 是一棵树，并且 不会有重复的边

// Assisted by BingChat
// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

// Solidity console.log
// https://hardhat.org/tutorial/debugging-with-hardhat-network#solidity--console.log
// import "hardhat/console.sol";

// 方法三：拓扑排序
contract MinHeightTrees_310 {
    function findMinHeightTrees(
        uint16 n,
        // fix: Type '[number, number][]' is not assignable to type 'readonly [readonly number[], readonly number[]]'.
        // uint16[][2] memory edges
        uint16[][] memory edges
    ) public pure returns (uint16[] memory) {
        require(n >= 0 && n <= 2e4, "0 <= n <= 2 * 10^4");
        if (n <= 0) {
            uint16[] memory x;
            return x;
        }
        if (n <= 1) {
            uint16[] memory x;
            x = append(x, 0);
            return x;
        }
        uint16[] memory deg = new uint16[](n);
        uint16[][] memory g = new uint16[][](n);
        for (uint16 i = 0; i < n; i++) {
            g[i] = new uint16[](0);
        }
        for (uint16 i = 0; i < edges.length; i++) {
            uint16 x = edges[i][0];
            uint16 y = edges[i][1];
            g[x] = append(g[x], y);
            g[y] = append(g[y], x);
            deg[x] += 1;
            deg[y] += 1;
        }

        uint16[] memory leaves = new uint16[](0);
        uint16 leafCount = 0;
        for (uint16 i = 0; i < n; i++) {
            if (deg[i] == 1) {
                leaves = append(leaves, i);
                leafCount += 1;
            }
        }

        uint16[] memory q;
        for (uint16 i = 0; i < n; i++) {
            if (deg[i] == 1) {
                q = append(q, i);
            }
        }
        uint16 remaining = n;
        while (remaining > 2) {
            remaining -= uint16(q.length); // q.length <= n
            uint16[] memory nxq;
            for (uint16 i = 0; i < q.length; i++) {
                uint16 x = q[i];
                // ** bug 1 - typo
                // for (uint16 y = 0; y < g[x].length; y++) {
                for (uint16 j = 0; j < g[x].length; j++) {
                    uint16 y = g[x][j];
                    deg[y] -= 1;
                    if (deg[y] == 1) {
                        nxq = append(nxq, y);
                    }
                }
            }
            q = nxq;
        }
        return q;
    }

    function append(
        uint16[] memory arr,
        uint16 x
    ) private pure returns (uint16[] memory) {
        uint16[] memory res = new uint16[](arr.length + 1);
        // ** bug 2 - forgot to fill the previus elements
        for (uint16 i = 0; i < arr.length; i++) {
            res[i] = arr[i];
        }
        res[arr.length] = x;
        return res;
    }
}
