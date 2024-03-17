import hre from "hardhat";
import { GetContractReturnType } from "@nomicfoundation/hardhat-viem/types";
import { expect } from "chai";

let cases = getCases();

describe("310. 最小高度树", () => {
  let solution: MinHeightTrees_310;
  before(async () => {
    let sol = await hre.viem.deployContract("MinHeightTrees_310");
    solution = sol;
  });

  for (let cs of cases) {
    let input = cs.slice(0, 2) as [number, [number, number][]];
    let want = cs[2] as number[];

    it(JSON.stringify(input), async () => {
      let ans = await solution.read.findMinHeightTrees(input);
      let a = [...ans].sort();
      let b = [...want].sort();
      // fix: AssertionError: expected [ +0 ] to equal [ +0 ]
      // expect(a).to.equal(b);
      expect(a).to.deep.equal(b);
    });
  }
});

function getCases() {
  return [
    [
      4,
      [
        [1, 0],
        [1, 2],
        [1, 3],
      ],
      [1],
    ],
    [
      6,
      [
        [3, 0],
        [3, 1],
        [3, 2],
        [3, 4],
        [5, 4],
      ],
      [3, 4],
    ],
    [2, [[0, 1]], [0, 1]],
    [1, [], [0]],
    [0, [], []],
  ];
}

type MinHeightTrees_310 = GetContractReturnType<
  [
    {
      inputs: [
        {
          internalType: "uint16";
          name: "n";
          type: "uint16";
        },
        {
          internalType: "uint16[][]";
          name: "edges";
          type: "uint16[][]";
        }
      ];
      name: "findMinHeightTrees";
      outputs: [
        {
          internalType: "uint16[]";
          name: "";
          type: "uint16[]";
        }
      ];
      stateMutability: "pure";
      type: "function";
    }
  ]
>;
