// 2789. 合并后数组中的最大元素
// https://leetcode.cn/problems/largest-element-in-an-array-after-merge-operations/
// 1 <= nums.length <= 10^5
// 1 <= nums[i] <= 10^6

// Asisted by 文心一言
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// Solidity console.log
// https://hardhat.org/tutorial/debugging-with-hardhat-network#solidity--console.log
import "hardhat/console.sol";

contract MaxArrayValue_2789 {
    uint32 constant MAX_SIZE = 1e6;
    // uint32[] private nums = new uint32[](MAX_SIZE);
    uint32[MAX_SIZE] private nums; // 定义一个包含1000个元素的固定大小数组
    uint32 private size;
    uint64 private result;

    // 设置数组值的函数
    function setArrayValues(uint32[] calldata _nums) external {
        // console.log("setArrayValues prev size", size);
        require(_nums.length <= 1000, "Array size must be within 1000");
        size = uint32(_nums.length);
        // console.log("setArrayValues new size", size);
        for (uint32 i = 0; i < size; i++) {
            nums[i] = _nums[i];
        }
        processArray(); // 设置完数组值后处理数组
    }

    // 处理数组的函数
    function processArray() private {
        // console.log("processArray size", size);
        if (size <= 0) {
            result = 0;
            return;
        }
        uint64 sum = nums[size - 1]; // 从最后一个元素开始
        if (size <= 1) {
            result = sum;
            return;
        }
        // ** bug - uint underflow
        // for (uint32 i = size - 2; i >= 0; i--) {
        uint32 i = size - 1;
        while (i > 0) {
            i--;
            if (nums[i] <= sum) {
                // unchecked {
                sum += nums[i];
                // }
            } else {
                sum = nums[i];
            }
        }
        result = sum;
    }

    // 获取处理后的结果
    function getResult() external view returns (uint64) {
        return result;
    }
}
