// Assisted by 文心一言
// SPDX-License-Identifier: MIT
import hre from "hardhat";
import { GetContractReturnType } from "@nomicfoundation/hardhat-viem/types";
import { expect } from "chai";

describe("2864. 最大二进制奇数", () => {
  let solution: GetContractReturnType<
    [
      {
        inputs: [
          {
            internalType: "string";
            name: "s";
            type: "string";
          }
        ];
        name: "maximumOddBinaryNumber";
        outputs: [
          {
            internalType: "string";
            name: "";
            type: "string";
          }
        ];
        stateMutability: "pure";
        type: "function";
      }
    ]
  >;

  before(async () => {
    solution = await hre.viem.deployContract("MaximumOddBinaryNumber_2864");
  });

  let cases = ["010", "001", "0101", "1001", "1", "1", "", ""];
  let us = 2;
  let cnt = cases.length / us;
  for (let i = 0; i < cnt; ++i) {
    let input = cases[i * us];
    let want = cases[i * us + 1];
    it(JSON.stringify(input), async () => {
      const ans = await solution.read.maximumOddBinaryNumber([input]);
      expect(ans).to.equal(want);
    });
  }
});
