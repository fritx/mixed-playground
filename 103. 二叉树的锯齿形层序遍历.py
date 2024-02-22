# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def zigzagLevelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        ans, q = [], [root]
        lv = 0
        while q:
            lv += 1
            ans.append([])
            n = len(q)
            sq, q = q, []
            for t in sq:
                if not t: continue
                q.append(t.left)
                q.append(t.right)
                ans[-1].append(t.val)
            if lv%2 == 0:
                ans[-1] = ans[-1][::-1]
        if ans and not len(ans[-1]):
            ans = ans[:-1]
        return ans
