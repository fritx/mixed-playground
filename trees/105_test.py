import json
from trees.tree import level_order_traversal
from utils.exec_module import exec_module

module = exec_module("105. 从前序与中序遍历序列构造二叉树.py")

def test():
    cases = json.loads("""[
        [3,9,20,15,7], [9,3,15,20,7], [3,9,20,null,null,15,7],
        [-1], [-1], [-1],
        [], [], []
    ]""")

    cnt = len(cases) // 3
    for i in range(cnt):
        data = cases[3*i : 3*i+2]
        want = cases[3*i+2]
        solution = module.Solution()
        root = solution.buildTree(data[0], data[1])

        # seq = serializeBinaryTree(root)
        # seq = serializeBinaryTree_2(root)
        seq = level_order_traversal(root)
        assert seq == want

        # wantRoot = buildTree(data[0], data[1])
        # assert is_same_tree(root, wantRoot)
