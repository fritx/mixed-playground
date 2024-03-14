// Assisted by 文心一言
// SPDX-License-Identifier: MIT
import hre from "hardhat";
import { GetContractReturnType } from "@nomicfoundation/hardhat-viem/types";
import { expect } from "chai";

describe("2129. 将标题首字母大写", () => {
  let solution: GetContractReturnType<
    [
      {
        inputs: [
          {
            internalType: "string";
            name: "input";
            type: "string";
          }
        ];
        name: "capitalizeTitle";
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
    solution = await hre.viem.deployContract("CapitalizeTheTitle_2129");
  });

  let cases = [
    "capiTalIze tHe titLe",
    "Capitalize The Title",
    "First leTTeR of EACH Word",
    "First Letter of Each Word",
    "i lOve leetcode",
    "i Love Leetcode",
    "hello world! this is a test.",
    "Hello World! This is a Test.",
    "hi there hi",
    "hi There hi",
    "",
    "",
    "   ",
    "   ",
    "hello-world! this_is? a#test.",
    "Hello-world! This_is? A#test.",
  ];
  let us = 2;
  let cnt = cases.length / us;
  for (let i = 0; i < cnt; ++i) {
    let input = cases[i * us];
    let want = cases[i * us + 1];
    it(JSON.stringify(input), async () => {
      let ans = await solution.read.capitalizeTitle([input]);
      expect(ans).to.equal(want);
    });
  }
});
