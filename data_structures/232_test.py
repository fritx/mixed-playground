import json
from utils.exec_module import exec_module

module = exec_module("232. 用栈实现队列.py")

def test():
    cases = json.loads("""[
["MyQueue","push","push","peek","pop","empty"], [[],[1],[2],[],[],[]], [null,null,null,1,1,false],
["MyQueue","push","push","push","push","pop","push","pop","pop","pop","pop"],
        [[],[1],[2],[3],[4],[],[5],[],[],[],[]],
        [null,null,null,null,null,1,null,2,3,4,5],
["MyQueue","push","push","pop","peek"], [[],[1],[2],[],[]], [null,null,null,1,2],
["MyQueue","push","push","pop","push","push","pop","peek"], [[],[1],[2],[],[3],[4],[],[]]
    ]""")
    us = 3
    cnt = len(cases) // us
    for Queue in [
        module.MyQueue,
        module.MyQueue2,
        module.MyQueue3,
    ]:
        for i in range(cnt):
            data = cases[us*i : us*(i+1)-1]
            want = cases[us*(i+1)-1]
            data[0] = data[0][1:]
            data[1] = data[1][1:]
            want = want[1:]
            n = len(data[0])
            ans = [None] * n
            q = Queue()
            for j, m in enumerate(data[0]):
                match m:
                    case "push":
                        q.push(*data[1][j])
                    case "pop":
                        ans[j] = q.pop()
                    case "peek":
                        ans[j] = q.peek()
                    case "empty":
                        ans[j] = q.empty()
            assert ans == want
