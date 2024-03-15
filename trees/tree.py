from collections import deque

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# @author 文心一言
def slice_to_binary_tree(nums):
    if not nums or nums[0] is None:
        return None

    root = TreeNode(nums[0])
    queue = [root]
    i = 1

    while queue and i < len(nums):
        node = queue.pop(0)

        if nums[i] is not None:
            node.left = TreeNode(nums[i])
            queue.append(node.left)
        i += 1

        if i < len(nums) and nums[i] is not None:
            node.right = TreeNode(nums[i])
            queue.append(node.right)
        i += 1

    return root

# @author 文心一言
def is_same_tree(p: TreeNode, q: TreeNode) -> bool:
    # 如果两个节点都为空，则它们是相等的
    if p is None and q is None:
        return True
    # 如果一个节点为空，而另一个不为空，则它们不相等
    if p is None or q is None:
        return False
    # 如果节点的值不相等，则它们不相等
    if p.val != q.val:
        return False
    # 递归地检查左子树和右子树是否相等
    return is_same_tree(p.left, q.left) and is_same_tree(p.right, q.right)

# @author 文心一言
def buildTree(preorder, postorder):
    if not preorder or not postorder:
        return None

    # 后序遍历的最后一个节点是根节点
    root_val = postorder[-1]
    root = TreeNode(root_val)

    # 在前序遍历中找到根节点的索引
    root_index = preorder.index(root_val)

    # 前序遍历中根节点左边的部分对应后序遍历中的左子树
    # 前序遍历中根节点右边的部分对应后序遍历中的右子树
    left_preorder = preorder[1:root_index + 1]
    right_preorder = preorder[root_index + 1:]
    left_postorder = postorder[:root_index]
    right_postorder = postorder[root_index:-1]

    # 递归构造左子树和右子树
    root.left = buildTree(left_preorder, left_postorder)
    root.right = buildTree(right_preorder, right_postorder)

    return root

# 如果左子树为空则用None占位，但是如果右子树为空则不需要None
# @author Bingo
def serializeBinaryTree(root):
    def dfs(node):
        if node is None:
            return []
        left = dfs(node.left)
        right = dfs(node.right)
        return [node.val] + ([None] if (not left and right) else left) + right
    return dfs(root)

# 还需要再调整一下，只要遇到空节点就用None占位，末尾的所有None可以省略
# @author Bingo
def serializeBinaryTree_2(root):
    def dfs(node):
        if node is None:
            return [None]
        left = dfs(node.left)
        right = dfs(node.right)
        return [node.val] + left + right
    result = dfs(root)
    while result and result[-1] is None:
        result.pop()
    return result

# 再调整一下，从前序遍历改为改为层序遍历，遇到空则放None
# @author Bingo
def level_order_traversal(root):
    if root is None:
        return []
    queue = deque([root])
    result = []
    while queue:
        node = queue.popleft()
        if node is None:
            result.append(None)
            continue
        result.append(node.val)
        # if node.left is not None or node.right is not None:
        queue.append(node.left)
        queue.append(node.right)
    while result and result[-1] is None:
        result.pop()
    return result
