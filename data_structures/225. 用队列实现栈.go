// 225. 用队列实现栈
// https://leetcode.cn/problems/implement-stack-using-queues/

package data_structures

// ===
// 一个队列 push-O(n),其他O(1)
// ===
type MyStack3 struct {
	queue []int
}

/** Initialize your data structure here. */
func Constructor3() (s MyStack3) {
	return
}

/** Push element x onto stack. */
func (s *MyStack3) Push(x int) {
	n := len(s.queue)
	s.queue = append(s.queue, x)
	for ; n > 0; n-- {
		s.queue = append(s.queue, s.queue[0])
		s.queue = s.queue[1:]
	}
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack3) Pop() int {
	v := s.queue[0]
	s.queue = s.queue[1:]
	return v
}

/** Get the top element. */
func (s *MyStack3) Top() int {
	return s.queue[0]
}

/** Returns whether the stack is empty. */
func (s *MyStack3) Empty() bool {
	return len(s.queue) == 0
}

// ===
// 两个队列 push-O(n),其他O(1)
// ===
type MyStack2 struct {
	queue1, queue2 []int
}

/** Initialize your data structure here. */
func Constructor2() (s MyStack2) {
	return
}

/** Push element x onto stack. */
func (s *MyStack2) Push(x int) {
	s.queue2 = append(s.queue2, x)
	for len(s.queue1) > 0 {
		s.queue2 = append(s.queue2, s.queue1[0])
		s.queue1 = s.queue1[1:]
	}
	s.queue1, s.queue2 = s.queue2, s.queue1
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack2) Pop() int {
	v := s.queue1[0]
	s.queue1 = s.queue1[1:]
	return v
}

/** Get the top element. */
func (s *MyStack2) Top() int {
	return s.queue1[0]
}

/** Returns whether the stack is empty. */
func (s *MyStack2) Empty() bool {
	return len(s.queue1) == 0
}

// ===
// 两个队列 push-O(1),其他O(n)
// ===
type MyStack struct {
	q1 []int
	q2 []int
}

func Constructor() MyStack {
	return MyStack{[]int{}, []int{}}
}

// Warning: receiver name should be a reflection of its identity;
// don't use generic names such as "this" or "self" (ST1006)go-staticcheck
func (this *MyStack) Push(x int) {
	if len(this.q1) > 0 && len(this.q2) > 0 {
		// opt.1
		// for len(this.q1) > 0 {
		//     transferOne(&this.q1, &this.q2)
		// }
		// for len(this.q2) > 0 {
		//     transferOne(&this.q2, &this.q1)
		// }
		// opt.2
		q := this.q1
		this.q1 = this.q2
		this.q2 = nil
		for len(q) > 0 {
			// ** bug 1 - typo
			// transferOne(*q, *this.q1)
			transferOne(&q, &this.q1)
		}
	}
	this.q1 = append(this.q1, x)
}
func (this *MyStack) Pop() int {
	x := this.Top()
	this.q1 = this.q1[1:]
	return x
}
func (this *MyStack) Top() int {
	if len(this.q1) == 0 {
		// opt.1
		// for len(this.q2) > 0 {
		//     transferOne(&this.q2, &this.q1)
		// }
		// opt.2
		this.q1 = this.q2
		this.q2 = nil
	}
	for len(this.q1) > 1 {
		transferOne(&this.q1, &this.q2)
	}
	return this.q1[0]
}
func (this *MyStack) Empty() bool {
	return len(this.q1) == 0 && len(this.q2) == 0
}
func transferOne(q1, q2 *[]int) {
	x := (*q1)[0]
	*q1 = (*q1)[1:]
	*q2 = append(*q2, x)
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
