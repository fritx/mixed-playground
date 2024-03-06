import importlib.util
import json, os

current_file = os.path.abspath(__file__)
current_dir = os.path.dirname(current_file)
src_file = os.path.join(current_dir, "2368. 受限条件下可到达节点的数目.py")
spec = importlib.util.spec_from_file_location("x", src_file)
module = importlib.util.module_from_spec(spec)
spec.loader.exec_module(module)

def test():
    cases = json.loads("""[
        7, [[0,1],[1,2],[3,1],[4,0],[0,5],[5,6]], [4,5], 4,
        7, [[0,1],[0,2],[0,5],[0,4],[3,2],[6,5]], [4,2,1], 3,
        1, [], [], 1,
        0, [], [], 0
    ]""")
    us = 4
    cnt = len(cases) // us
    for Solution in [
        module.Solution,
        module.Solution2,
        module.Solution3,
    ]:
        s = Solution()
        sname = s.__class__.__name__
        for i in range(cnt):
            data = cases[us*i:us*i+3]
            want = cases[us*i+3]
            try:
                ans = s.reachableNodes(*data)
            except:
                assert False, sname
            assert ans == want, sname


