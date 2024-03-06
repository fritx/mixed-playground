import json
from tree import slice_to_binary_tree
from utils.exec_module import exec_module

module = exec_module("102. 二叉树的层序遍历.py")

def test():
    cases = json.loads("""[
        [3,9,20,null,null,15,7], [[3],[9,20],[15,7]],
        [1], [[1]],
        [], []
    ]""")

    cnt = len(cases) // 2
    for i in range(cnt):
        data = cases[2*i]
        want = cases[2*i+1]
        solution = module.Solution()
        root = slice_to_binary_tree(data)
        ans = solution.levelOrder(root)
        assert ans == want
