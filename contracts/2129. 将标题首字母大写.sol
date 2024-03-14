// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// Assisted by 文心一言
contract CapitalizeTheTitle_2129 {
    // 将单个字符转换为小写（如果它是大写字母）
    function toLowerCase(uint8 char) internal pure returns (uint8) {
        if (char >= 65 && char <= 90) {
            // 如果是大写字母
            return char + 32; // 转换为小写
        }
        return char;
    }

    // 将单个字符转换为大写（如果它是小写字母）
    function toUpperCase(uint8 char) internal pure returns (uint8) {
        if (char >= 97 && char <= 122) {
            // 如果是小写字母
            return char - 32; // 转换为大写
        }
        return char;
    }

    // 转换字符串中每个单词的首字母为大写，其他字母为小写
    function capitalizeTitle(
        string memory input
    ) public pure returns (string memory) {
        bytes memory inputBytes = bytes(input);
        uint256 length = inputBytes.length;
        bool isWordStart = true; // 标记是否为单词的开始

        // 遍历字符串中的每个字符
        for (uint256 i = 0; i < length; i++) {
            uint8 char = uint8(inputBytes[i]);

            // 如果是空格，则下一个字符是单词的开始
            if (char == 32) {
                isWordStart = true;
            } else {
                // 转换大小写
                if (isWordStart) {
                    // 如果是单词的开始，并且单词长度大于2，则转换为大写
                    if (
                        i + 2 < length &&
                        uint8(inputBytes[i + 1]) != 32 &&
                        uint8(inputBytes[i + 2]) != 32
                    ) {
                        inputBytes[i] = bytes1(toUpperCase(char));
                    } else {
                        // 如果单词长度小于等于2，则全转为小写
                        inputBytes[i] = bytes1(toLowerCase(char));
                    }
                    isWordStart = false;
                } else {
                    // 不是单词的开始，则转换为小写
                    inputBytes[i] = bytes1(toLowerCase(char));
                }
            }
        }

        // 将修改后的字节数组转换回字符串
        return string(inputBytes);
    }
}
