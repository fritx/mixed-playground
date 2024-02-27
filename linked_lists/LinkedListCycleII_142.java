// 142. 环形链表 II.java
// https://leetcode.cn/problems/linked-list-cycle-ii/
package linked_lists;

import java.util.HashSet;
import java.util.Set;

// 测试用例
// [3,2,0,-4], 1, "tail connects to node index 1"
// [1,2], 0, "tail connects to node index 0"
// [1], -1, "no cycle"

/**
 * Definition for singly-linked list.
 * class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) {
 *         val = x;
 *         next = null;
 *     }
 * }
 */
public class LinkedListCycleII_142 {
    public ListNode detectCycle(ListNode head) {
        Set<ListNode> seen = new HashSet<ListNode>();
        while (head != null) {
            if (seen.contains(head)) {
                return head;
            }
            seen.add(head);
            head = head.next;
        }
        return null;
    }
}
