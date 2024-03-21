// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// Assisted by 文心一言
contract FrequencyTracker_2671 {
    // 使用 mapping 存储每个数字的频率
    mapping(int256 => uint256) private freq;
    // 使用 mapping 存储每个频率的计数
    mapping(uint256 => uint256) private freqCnt;

    // 添加数字到频率跟踪器
    function add(int256 number) public {
        uint256 prevFreq = freq[number];
        // 减少旧频率的计数
        if (prevFreq > 0) {
            freqCnt[prevFreq] -= 1;
        }
        // 增加数字的频率
        freq[number] += 1;
        uint256 newFreq = freq[number];
        // 增加新频率的计数
        freqCnt[newFreq] += 1;
    }

    // 从频率跟踪器中删除一个数字（如果它存在）
    function deleteOne(int256 number) public {
        uint256 prevFreq = freq[number];
        if (prevFreq == 0) {
            return; // 数字的频率已经是 0，无需操作
        }
        // 减少旧频率的计数
        freqCnt[prevFreq] -= 1;
        // 减少数字的频率
        freq[number] -= 1;
        uint256 newFreq = freq[number];
        // 如果新频率大于 0，则增加它的计数
        if (newFreq > 0) {
            freqCnt[newFreq] += 1;
        }
    }

    // 检查是否存在具有指定频率的数字
    function hasFrequency(uint256 _frequency) public view returns (bool) {
        return freqCnt[_frequency] > 0;
    }
}
