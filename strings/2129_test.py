import json
from utils.exec_module import exec_module

module = exec_module("2129. 将标题首字母大写.py")

def test():
    cases = json.loads("""[
        "capiTalIze tHe titLe", "Capitalize The Title",
		"First leTTeR of EACH Word", "First Letter of Each Word",
		"i lOve leetcode", "i Love Leetcode"
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
            ans = solution.capitalizeTitle(input)
            assert ans == want
