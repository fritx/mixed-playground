# 297. 二叉树的序列化与反序列化
# https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/

from trees.tree import TreeNode

# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

# 方法一：深度优先搜索
# 方法二：括号表示编码 + 递归下降解码
# todo

# 自己解法
class Codec:
    sep = ","
    null = "null"

    def serialize(self, root):
        """Encodes a tree to a single string.
        :type root: TreeNode
        :rtype: str
        """
        if not root: return "[]"
        q = [root]
        ans = []
        while q:
            n = q[0]
            q = q[1:]
            if not n:
                ans.append(None)
                continue
            ans.append(n.val)
            q.append(n.left)
            q.append(n.right)
        s = self.sep.join(map(lambda x : self.null if x is None else str(x), ans))
        return "[" + s + "]"


    def deserialize(self, data):
        """Decodes your encoded data to tree.
        :type data: str
        :rtype: TreeNode
        """
        assert data[0] == "[" and data[-1] == "]"
        data = data[1:-1]
        # ** bug - 常见陷阱！ "".split(",") == "" 不符预期
        if not data: return None
        arr = data.split(self.sep)
        if not arr: return None
        if arr[0] == self.null : return None
        i = 0
        n = len(arr)
        # ** bug - 根节点也要确保str->int
        # 题目OJ能通过 但本地测试用例不通过
        root = TreeNode(arr[i])
        root = TreeNode(int(arr[i]))
        q = [root]
        while q:
            c = q[0]
            q = q[1:]
            i += 1
            if i < n and arr[i] != self.null:
                c.left = TreeNode(int(arr[i]))
                q.append(c.left)
            i += 1
            if i < n and arr[i] != self.null:
                c.right = TreeNode(int(arr[i]))
                q.append(c.right)
            # ** bug - 应放入if分支中
            # q.append(c.left)
            # q.append(c.right)
        return root

# Your Codec object will be instantiated and called as such:
# ser = Codec()
# deser = Codec()
# ans = deser.deserialize(ser.serialize(root))
