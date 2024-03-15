from typing import Optional, List
from trees.tree import TreeNode

# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

# 递归
class Solution:
    def buildTree(self, inorder: List[int], postorder: List[int]) -> Optional[TreeNode]:

        # ** bug 1
        if not inorder: return None

        rv=postorder[-1]
        r=TreeNode(rv)
        idx=inorder.index(rv)
        ll=idx
        #rl=len(inorder)-idx-1
        r.left=self.buildTree(inorder[:idx], postorder[:ll])

        # ** bug 2              #r.right=self.buildTree(inorder[idx+1:], postorder[ll:1])
        r.right=self.buildTree(inorder[idx+1:], postorder[ll:-1])

        return r
