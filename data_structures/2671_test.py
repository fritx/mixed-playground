import json
from utils.exec_module import exec_module

module = exec_module("2671. 频率跟踪器.py")

def test():
    cases = json.loads("""[
        [["FrequencyTracker","add","add","hasFrequency"], [[],[3],[3],[2]], [null,null,null,true]],
        [["FrequencyTracker","add","deleteOne","hasFrequency"], [[],[1],[1],[1]], [null,null,null,false]],
        [["FrequencyTracker","hasFrequency","add","hasFrequency"], [[],[2],[3],[1]], [null,false,null,true]],
        [[], [], []]
    ]""")
    us = 3
    for FrequencyTracker in [
        module.FrequencyTracker,
    ]:
        for case in cases:
            data = case[ : us-1]
            want = case[us-1]
            data[0] = data[0][1:]
            data[1] = data[1][1:]
            want = want[1:]
            n = len(data[0])
            ans = [None] * n
            t = FrequencyTracker()
            for j, m in enumerate(iterable=data[0]):
                match m:
                    case "add":
                        ans[j] = t.add(*data[1][j])
                    case "deleteOne":
                        ans[j] = t.deleteOne(*data[1][j])
                    case "hasFrequency":
                        ans[j] = t.hasFrequency(*data[1][j])
            assert ans == want
