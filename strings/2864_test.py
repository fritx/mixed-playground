import json
from utils.exec_module import exec_module

module = exec_module("2864. 最大二进制奇数.py")

def test():
    cases = json.loads("""[
        "010", "001",
		"0101", "1001",
		"1", "1",
		"", ""
    ]""")
    us = 2
    cnt = len(cases) // us
    for Solution in [
        module.Solution,
    ]:
        for i in range(cnt):
            input = cases[us*i]
            want = cases[us*i+1]
            solution = Solution()
            ans = solution.maximumOddBinaryNumber(input)
            assert ans == want
