package me.fritx.trees;

import java.util.LinkedList;
import java.util.Queue;

public class TreeUtils {

    // public static TreeNode binaryTreeFromArray(int[] nums) {
    public static TreeNode binaryTreeFromArray(Integer[] nums) {
        if (nums == null || nums.length == 0 || nums[0] == -1) { // Assuming -1 represents null in Java
            return null;
        }

        TreeNode root = new TreeNode(nums[0]);
        Queue<TreeNode> queue = new LinkedList<>();
        queue.add(root);
        int i = 1;

        while (!queue.isEmpty() && i < nums.length) {
            TreeNode node = queue.poll();

            // if (nums[i] != -1) { // Assuming -1 represents null in Java
            if (nums[i] != null) {
                node.left = new TreeNode(nums[i]);
                queue.add(node.left);
            }
            i++;

            // if (i < nums.length && nums[i] != -1) { // Assuming -1 represents null in
            // Java
            if (i < nums.length && nums[i] != null) {
                node.right = new TreeNode(nums[i]);
                queue.add(node.right);
            }
            i++;
        }

        return root;
    }

    public static boolean isSameTree(TreeNode p, TreeNode q) {
        // 如果两个节点都为空，则它们是相等的
        if (p == null && q == null) {
            return true;
        }
        // 如果一个节点为空，而另一个不为空，则它们不相等
        if (p == null || q == null) {
            return false;
        }
        // 如果节点的值不相等，则它们不相等
        if (p.val != q.val) {
            return false;
        }
        // 递归地检查左子树和右子树是否相等
        return isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
    }
}
