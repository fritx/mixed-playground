# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        ans = []

        q = [root]
        nq = []

        while q or nq:
            ans.append([])
            while q:
                n = q[0]
                q = q[1:]
                if not n: continue
                ans[-1].append(n.val)
                nq.append(n.left)
                nq.append(n.right)
            q = nq
            nq = []

        if ans and not ans[-1]:
            ans = ans[:-1]
        return ans
