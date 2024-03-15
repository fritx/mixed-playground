// Assisted by 文心一言
// SPDX-License-Identifier: MIT
import hre from "hardhat";
import { GetContractReturnType } from "@nomicfoundation/hardhat-viem/types";
import { expect } from "chai";

describe("2789. 合并后数组中的最大元素", () => {
  let solution: GetContractReturnType<
    [
      {
        inputs: [];
        name: "getResult";
        outputs: [
          {
            internalType: "uint64";
            name: "";
            type: "uint64";
          }
        ];
        stateMutability: "view";
        type: "function";
      },
      {
        inputs: [
          {
            internalType: "uint32[]";
            name: "_nums";
            type: "uint32[]";
          }
        ];
        name: "setArrayValues";
        outputs: [];
        stateMutability: "nonpayable";
        type: "function";
      }
    ]
  >;
  before(async () => {
    let sol = await hre.viem.deployContract("MaxArrayValue_2789");
    solution = sol;
  });

  let cases = [[2, 3, 7, 9, 3], 21, [5, 3, 3], 11, [91, 50], 91, [3], 3, [], 0];
  let us = 2;
  let cnt = cases.length / us;

  for (let i = 0; i < cnt; ++i) {
    let input = cases[i * us] as number[];

    // BigInt required for `uint64`
    // fix: expected 21n to equal 21
    let want = BigInt(cases[i * us + 1] as number);

    it(JSON.stringify(input), async () => {
      await solution.write.setArrayValues([input]);
      let ans = await solution.read.getResult();
      expect(ans).to.equal(want);
    });
  }
});
