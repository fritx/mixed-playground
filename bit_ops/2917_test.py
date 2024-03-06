import importlib.util
import json, os

# fix: FileNotFoundError: [Errno 2] No such file or directory: 'xxx/mixed-playground/xxx.py'
current_file = os.path.abspath(__file__)
current_dir = os.path.dirname(current_file)
src_file = os.path.join(current_dir, "2917. 找出数组中的 K-or 值.py")

spec = importlib.util.spec_from_file_location("x", src_file)
module = importlib.util.module_from_spec(spec)
spec.loader.exec_module(module)

def test():
    cases = json.loads("""[
        [7,12,9,8,9,15], 4, 9,
        [2,12,1,11,4,5], 6, 0,
        [10,8,5,9,11,6,8], 1, 15
    ]""")
    us = 3
    cnt = len(cases) // us
    for Solution in [
        module.Solution,
        module.Solution2,
        module.Solution3,
    ]:
        s = Solution()
        sname = s.__class__.__name__
        for i in range(cnt):
            data = cases[us*i:us*(i+1)-1]
            want = cases[us*(i+1)-1]
            try:
                ans = s.findKOr(*data)
            except:
                assert False, sname
            assert ans == want, sname
