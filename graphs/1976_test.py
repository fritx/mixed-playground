import importlib.util
import json, os

current_file = os.path.abspath(__file__)
current_dir = os.path.dirname(current_file)
src_file = os.path.join(current_dir, "1976. 到达目的地的方案数.py")
spec = importlib.util.spec_from_file_location("x", src_file)
module = importlib.util.module_from_spec(spec)
spec.loader.exec_module(module)

def test_reachableNodes():
    cases = json.loads("""[
        7, [[0,6,7],[0,1,2],[1,2,3],[1,3,3],[6,3,3],[3,5,1],[6,5,1],[2,5,1],[0,4,5],[4,6,2]], 4,
        2, [[1,0,10]], 1,
        2, [], 0
    ]""")
    us = 3
    cnt = len(cases) // us
    for Solution in [
        module.Solution,
    ]:
        s = Solution()
        sname = s.__class__.__name__
        for i in range(cnt):
            data = cases[us*i:us*(i+1)-1]
            want = cases[us*(i+1)-1]
            try:
                ans = s.countPaths(*data)
            except:
                assert False, sname
            assert ans == want, sname
