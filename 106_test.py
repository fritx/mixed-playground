from tree import level_order_traversal
import importlib.util
import json

spec = importlib.util.spec_from_file_location("x", "106. 从中序与后序遍历序列构造二叉树.py")
module = importlib.util.module_from_spec(spec)
spec.loader.exec_module(module)

def test_buildTree():
    cases = json.loads("""[
        [9,3,15,20,7], [9,15,7,20,3], [3,9,20,null,null,15,7],
        [-1], [-1], [-1],
        [], [], []
    ]""")

    cnt = len(cases) // 3
    for i in range(cnt):
        data = cases[3*i : 3*i+2]
        want = cases[3*i+2]
        solution = module.Solution()
        root = solution.buildTree(data[0], data[1])
        seq = level_order_traversal(root)
        assert seq == want
