// 141. 环形链表
// https://leetcode.cn/problems/linked-list-cycle/
package linked_lists;

import java.util.HashSet;
import java.util.Set;

// 测试用例
// [3,2,0,-4], 1, true
// [1,2], 0, true
// [1], -1, false
// [1,2], -1, false

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
public class LinkedListCycle_141 {
    public boolean hasCycle(ListNode head) {
        Set<ListNode> seen = new HashSet<ListNode>();
        while (head != null) {
            if (seen.contains(head)) {
                return true;
            }
            // ** bug 1
            // head = head.next;
            // seen.add(head);
            seen.add(head);
            head = head.next;
        }
        return false;
    }
}
