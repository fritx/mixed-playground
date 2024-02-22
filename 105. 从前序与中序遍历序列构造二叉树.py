# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

# 递归
# O(n), O(h)
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if len(preorder) == 0:
            return None
        root_val = preorder[0]
        idx = inorder.index(root_val)
        left_inorder = inorder[: idx]
        left_preorder = preorder[1: len(left_inorder)+1]
        right_inorder = inorder[idx+1 :]
        right_preorder = preorder[len(left_inorder)+1 :]
        root = TreeNode(root_val)
        root.left = self.buildTree(left_preorder, left_inorder)
        root.right = self.buildTree(right_preorder, right_inorder)
        return root
