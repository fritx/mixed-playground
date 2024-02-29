package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 迭代 / O(n),O(1)
func reverseList_it(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// ** bug 1 - 需要引入哨兵节点
	// newHead := head
	// for head != nil && head.Next != nil {
	//     newHead := head.Next.Next
	//     head.Next.Next = head
	//     newHead.Next = head.Next
	// }
	// return newHead
	// ** bug 2 - use of untyped nil in assignment
	// left, right := nil, head
	var pre, curr *ListNode = nil, head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre, curr = curr, next
	}
	return pre
}

// 递归 / O(n),O(n)
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := reverseList(head.Next)
	head.Next.Next = head
	// ** bug 1 - missing
	head.Next = nil
	return tail
}
