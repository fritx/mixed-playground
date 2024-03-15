import hre from "hardhat";
import { GetContractReturnType } from "@nomicfoundation/hardhat-viem/types";
import { expect } from "chai";

let cases = getCases();

describe("2312. 卖木头块", () => {
  let solution: SellingWood_2312;
  before(async () => {
    let sol = await hre.viem.deployContract("SellingWood_2312");
    solution = sol;
  });

  for (let cs of cases) {
    let input = cs.slice(0, 3) as [number, number, number[][]];

    // BigInt required for `uint64`
    let want = BigInt(cs[3] as number);

    it(JSON.stringify(input), async () => {
      let ans = await solution.read.sellingWood(input);
      expect(ans).to.equal(want);
    });
  }
});

function getCases() {
  return [
    [
      3,
      5,
      [
        [1, 4, 2],
        [2, 2, 7],
        [2, 1, 3],
      ],
      19,
    ],
    [
      4,
      6,
      [
        [3, 2, 10],
        [1, 4, 2],
        [4, 1, 3],
      ],
      32,
    ],
    [7, 7, [], 0],
    [0, 0, [[1, 2, 3]], 0],
    [0, 0, [], 0],
  ];
}

type SellingWood_2312 = GetContractReturnType<
  [
    {
      inputs: [
        {
          internalType: "uint8";
          name: "m";
          type: "uint8";
        },
        {
          internalType: "uint8";
          name: "n";
          type: "uint8";
        },
        {
          internalType: "uint32[][]";
          name: "prices";
          type: "uint32[][]";
        }
      ];
      name: "sellingWood";
      outputs: [
        {
          internalType: "int64";
          name: "";
          type: "int64";
        }
      ];
      stateMutability: "pure";
      type: "function";
    }
  ]
>;
