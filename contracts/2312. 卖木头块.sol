// SPDX-License-Identifier: MIT
// pragma solidity ^0.8.0; // 文心一言
pragma solidity >=0.4.22 <0.9.0; // BingChat

// Assisted by BingChat
contract SellingWood_2312 {
    function sellingWood(
        uint8 m,
        uint8 n,
        // fix: AbiEncodingArrayLengthMismatchError: ABI encoding array length mismatch for type uint32[][3].
        //      Expected length: 3
        //      Given length: 0
        // uint32[][3] memory prices
        uint32[][] memory prices
    ) public pure returns (int64) {
        // 1 <= m, n <= 200
        // 1 <= prices.length <= 2 * 10^4
        // prices[i].length == 3
        // 1 <= hi <= m
        // 1 <= wi <= n
        // 1 <= price[i] <= 10^6
        require(m >= 0 && m <= 200 && n >= 0 && n <= 200, "0 <= m, n <= 200");
        require(
            prices.length >= 0 && prices.length <= 2 * 1e4,
            "0 <= prices.length <= 2 * 10^4"
        );
        int64[201][201] memory value;
        int64[201][201] memory dp;

        for (uint16 i = 0; i < prices.length; i++) {
            // require(prices[i][0] >= 0 && prices[i][0] <= m, "0 <= hi <= m");
            // require(prices[i][1] >= 0 && prices[i][1] <= n, "0 <= wi <= n");
            require(
                prices[i][2] >= 0 && prices[i][2] <= 1e6,
                "0 <= price[i] <= 10^6"
            );
            // fix: Explicit type conversion not allowed from "uint16" to "int64".solidity(9640)
            // value[prices[i][0]][prices[i][1]] = int64(prices[i][2]);
            value[prices[i][0]][prices[i][1]] = int64(uint64(prices[i][2]));
        }
        for (uint8 x = 1; x <= m; x++) {
            for (uint8 y = 1; y <= n; y++) {
                dp[x][y] = value[x][y];
                for (uint8 i = 1; i < x; i++) {
                    dp[x][y] = max(dp[x][y], dp[i][y] + dp[x - i][y]);
                }
                for (uint8 j = 1; j < y; j++) {
                    dp[x][y] = max(dp[x][y], dp[x][j] + dp[x][y - j]);
                }
            }
        }
        return dp[m][n];
    }

    function max(int64 a, int64 b) private pure returns (int64) {
        return a > b ? a : b;
    }
}
