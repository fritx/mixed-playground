import json
from tree import slice_to_binary_tree, is_same_tree
from utils.exec_module import exec_module

module = exec_module("297. 二叉树的序列化与反序列化.py")

def test():
    cases = json.loads(
        """[
        [1,2,3,null,null,4,5], [1,2,3,null,null,4,5],
        [1,null,2], [1,null,2],
        [1,2], [1,2],
        [], []
    ]""")
    us = 2
    cnt = len(cases) // us
    for i in range(cnt):
        data = cases[us * i]
        want = cases[us * i + 1]
        root = slice_to_binary_tree(data)
        Codec = module.Codec
        ser = Codec()
        deser = Codec()
        ansRoot = deser.deserialize(ser.serialize(root))
        # s = level_order_traversal(ansRoot)
        # assert s == want
        wantRoot = slice_to_binary_tree(want)
        assert is_same_tree(ansRoot, wantRoot)
