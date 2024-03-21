import hre from "hardhat";
import { GetContractReturnType } from "@nomicfoundation/hardhat-viem/types";
import { expect } from "chai";

let cases = getCases();

describe("2671. 频率跟踪器", () => {
  let ft: FrequencyTracker_2671;

  // Each ft is used as a separate instance
  // before(async () => {
  beforeEach(async () => {
    let t = await hre.viem.deployContract("FrequencyTracker_2671");
    ft = t;
  });

  for (let cs of cases) {
    let methods = cs[0].slice(1) as string[];
    let args = (cs[1].slice(1) as [number][]).map<[bigint]>(([v]) => [
      BigInt(v),
    ]);
    let want = cs[2].slice(1) as (null | boolean)[];

    // fix: Do not know how to serialize a BigInt
    // https://github.com/GoogleChromeLabs/jsbi/issues/30#issuecomment-521460510
    it(
      JSON.stringify(
        [methods, args],
        (_, v) => (typeof v === "bigint" ? `${v}n` : v) // return everything else unchanged
      ),
      async () => {
        let n = methods.length;
        let ans = new Array(n);
        ans.fill(null);
        for (let i = 0; i < n; i++) {
          switch (methods[i]) {
            case "add":
              await ft.write.add(args[i]);
              break;
            case "deleteOne":
              await ft.write.deleteOne(args[i]);
              break;
            case "hasFrequency":
              ans[i] = await ft.read.hasFrequency(args[i]);
              break;
          }
        }
        expect(ans).to.deep.equal(want);
      }
    );
  }
});

function getCases() {
  return [
    [
      ["FrequencyTracker", "add", "add", "hasFrequency"],
      [[], [3], [3], [2]],
      [null, null, null, true],
    ],
    [
      ["FrequencyTracker", "add", "deleteOne", "hasFrequency"],
      [[], [1], [1], [1]],
      [null, null, null, false],
    ],
    [
      ["FrequencyTracker", "hasFrequency", "add", "hasFrequency"],
      [[], [2], [3], [1]],
      [null, false, null, true],
    ],
    [[], [], []],
  ];
}

type FrequencyTracker_2671 = GetContractReturnType<
  [
    {
      inputs: [
        {
          internalType: "int256";
          name: "number";
          type: "int256";
        }
      ];
      name: "add";
      outputs: [];
      stateMutability: "nonpayable";
      type: "function";
    },
    {
      inputs: [
        {
          internalType: "int256";
          name: "number";
          type: "int256";
        }
      ];
      name: "deleteOne";
      outputs: [];
      stateMutability: "nonpayable";
      type: "function";
    },
    {
      inputs: [
        {
          internalType: "uint256";
          name: "_frequency";
          type: "uint256";
        }
      ];
      name: "hasFrequency";
      outputs: [
        {
          internalType: "bool";
          name: "";
          type: "bool";
        }
      ];
      stateMutability: "view";
      type: "function";
    }
  ]
>;
