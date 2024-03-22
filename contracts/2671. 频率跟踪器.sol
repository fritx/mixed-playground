// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// 1 <= number <= 10^5
// 1 <= frequency <= 10^5
// 最多调用 add、deleteOne 和 hasFrequency 共计 2 * 10^5 次

// Assisted by 文心一言
contract FrequencyTracker_2671 {
    uint32 private constant maxN = 1e5;

    // 使用 mapping 存储每个数字的频率
    mapping(uint32 => uint32) private freq;
    // 使用 mapping 存储每个频率的计数
    mapping(uint32 => uint32) private freqCnt;

    // 添加数字到频率跟踪器
    function add(uint32 number) public {
        require(number >= 1 && number <= maxN, "1 <= number <= 10^5");

        uint32 prevFreq = freq[number];
        // 减少旧频率的计数
        if (prevFreq > 0) {
            freqCnt[prevFreq] -= 1;
        }

        // 增加数字的频率
        require(
            freq[number] >= 0 && freq[number] < maxN,
            "1 <= frequency <= 10^5"
        );
        freq[number] += 1;
        uint32 newFreq = freq[number];

        // 增加新频率的计数
        freqCnt[newFreq] += 1;
    }

    // 从频率跟踪器中删除一个数字（如果它存在）
    function deleteOne(uint32 number) public {
        require(number >= 1 && number <= maxN, "1 <= number <= 10^5");

        uint32 prevFreq = freq[number];
        if (prevFreq == 0) {
            return; // 数字的频率已经是 0，无需操作
        }
        // 减少旧频率的计数
        freqCnt[prevFreq] -= 1;
        // 减少数字的频率
        freq[number] -= 1;
        uint32 newFreq = freq[number];
        // 如果新频率大于 0，则增加它的计数
        if (newFreq > 0) {
            freqCnt[newFreq] += 1;
        }
    }

    // 检查是否存在具有指定频率的数字
    function hasFrequency(uint32 _frequency) public view returns (bool) {
        return freqCnt[_frequency] > 0;
    }
}
