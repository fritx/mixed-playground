// Assisted by 文心一言
// SPDX-License-Identifier: MIT
import hre from "hardhat";
import { GetContractReturnType } from "@nomicfoundation/hardhat-viem/types";
import { expect } from "chai";

describe("CapitalizeTheTitle_2129", () => {
  let solution: GetContractReturnType<[{
    inputs: [{
      internalType: "string";
      name: "input";
      type: "string";
    }];
    name: "capitalizeTitle";
    outputs: [{
      internalType: "string";
      name: "";
      type: "string";
    }];
    stateMutability: "pure";
    type: "function";
  }]>;

  before(async () => {
    solution = await hre.viem.deployContract("CapitalizeTheTitle_2129");
  });

  it("should capitalize the first letter of each word", async () => {
    const input = "hello world! this is a test.";
    const expectedOutput = "Hello World! This is a Test.";
    const output = await solution.read.capitalizeTitle([input]);
    expect(output).to.equal(expectedOutput);
  });

  it("should capitalize the first letter of each word with short words", async () => {
    const input = "hi there hi";
    const expectedOutput = "hi There hi";
    const output = await solution.read.capitalizeTitle([input]);
    expect(output).to.equal(expectedOutput);
  });

  it("should handle empty string", async () => {
    const input = "";
    const expectedOutput = "";
    const output = await solution.read.capitalizeTitle([input]);
    expect(output).to.equal(expectedOutput);
  });

  it("should handle string with only spaces", async () => {
    const input = "   ";
    const expectedOutput = "   ";
    const output = await solution.read.capitalizeTitle([input]);
    expect(output).to.equal(expectedOutput);
  });

  it("should handle string with special characters", async () => {
    const input = "hello-world! this_is? a#test.";
    const expectedOutput = "Hello-world! This_is? A#test.";
    const output = await solution.read.capitalizeTitle([input]);
    expect(output).to.equal(expectedOutput);
  });
});
