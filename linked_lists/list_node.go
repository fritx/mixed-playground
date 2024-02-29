package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToLinkedList(data []int) *ListNode {
	head := &ListNode{}
	curr := head
	for _, v := range data {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head.Next
}

func linkedListToSlice(head *ListNode) []int {
	s := make([]int, 0)
	for head != nil {
		s = append(s, head.Val)
		head = head.Next
	}
	return s
}
