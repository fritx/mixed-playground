// Assisted by 文心一言
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MaximumOddBinaryNumber_2864 {
    function maximumOddBinaryNumber(
        string memory s
    ) public pure returns (string memory) {
        // 将字符串转换为字节数组
        bytes memory bytesS = bytes(s);
        uint256 n = bytesS.length;
        uint256 oneCount = 0;

        // fix: Details: VM Exception while processing transaction: reverted with panic code 0x11
        //      (Arithmetic operation overflowed outside of an unchecked block)
        // 因为后面用到了 (uint) n-1 所以溢出了
        // 题目OJ能通过 但本地测试用例不通过
        if (n == 0) {
            return "";
        }

        // 统计'1'的数量
        for (uint256 i = 0; i < n; i++) {
            if (bytesS[i] == "1") {
                oneCount++;
            }
        }

        // 创建一个足够大的内存数组来存储结果
        bytes memory ans = new bytes(n);

        // 填充数组
        for (uint256 i = 0; i < n - 1; i++) {
            if (i < oneCount - 1) {
                ans[i] = "1";
            } else {
                ans[i] = "0";
            }
        }

        // 最后一位固定为'1'
        if (n > 0) {
            ans[n - 1] = "1";
        }

        return string(ans);
    }
}
