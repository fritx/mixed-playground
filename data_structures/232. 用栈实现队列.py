# 232. 用栈实现队列
# https://leetcode.cn/problems/implement-queue-using-stacks/

# 测试用例
# - ["MyQueue","push","push","peek","pop","empty"], [[],[1],[2],[],[],[]], [null,null,null,1,1,false]
# - ["MyQueue","push","push","push","push","pop","push","pop","pop","pop","pop"],
# [[],[1],[2],[3],[4],[],[5],[],[],[],[]],
# [null,null,null,null,null,1,null,2,3,4,5]
# - ["MyQueue","push","push","pop","peek"], [[],[1],[2],[],[]], [null,null,null,1,2]
# - ["MyQueue","push","push","pop","push","push","pop","peek"], [[],[1],[2],[],[3],[4],[],[]]

# 方法一：双栈
# 每个操作的均摊时间复杂度是 O(1)
class MyQueue3:
    def __init__(self):
        self.instk = []
        self.outstk = []

    def push(self, x: int) -> None:
        self.instk.append(x)

    def in2out(self):
        while self.instk:
            self.outstk.append(self.instk.pop())

    def pop(self) -> int:
        if not self.outstk:
            self.in2out()
        return self.outstk.pop()

    def peek(self) -> int:
        if not self.outstk:
            self.in2out()
        return self.outstk[-1]

    def empty(self) -> bool:
        return not self.instk and not self.outstk

# 路人 双栈
class MyQueue:
    def __init__(self):
        self.data = []  # 数据栈
        self.stack = [] # 辅助栈

    def push(self, x: int) -> None:
        # 先把数据倒腾出去: data -> stack
        while self.data:
            self.stack.append(self.data.pop())

        # 新元素入队
        self.data.append(x)

        # 再把数据倒腾回来: stack -> data
        while self.stack:
            self.data.append(self.stack.pop())

    def pop(self) -> int:
        return self.data.pop()

    def peek(self) -> int:
        return self.data[-1]

    def empty(self) -> bool:
        return len(self.data) == 0

# 自己解法 双栈
# 每个操作的均摊时间复杂度可以认为是 O(1)，假设 push 操作远多于 pop 和 peek 操作。
# 然而，需要注意的是，这个实现并不是最优的，因为 pop 和 peek 操作在最坏情况下的时间复杂度是 O(n)。
# 一个更高效的实现可以使用一个栈来存储元素，并在 pop 和 peek 操作时将元素逐个移动到另一个栈中，从而保持时间复杂度为 O(1)。
class MyQueue2:
    def __init__(self):
        self.s1 = []
        self.s2 = []

    def push(self, x: int) -> None:
        while self.s2:
            # ** bug 2 - 再次掉入同一陷阱，需小心：python中，局部声明变量，即使是块/分支内的语句，
            # 同名变量会直接在整个函数内部生效，覆盖参数变量，语法及运行均无警告，导致不符预期/可能长时间debug的坑，需注意！
            # x = self.s2.pop()
            # self.s1.append(x)
            y = self.s2.pop()
            self.s1.append(y)
        self.s1.append(x)

    def pop(self) -> int:
        # ** bug 3.1 - 忽略了s1为空的情况
        if not self.s1:
            return self.s2.pop()
        while len(self.s1) > 1:
            x = self.s1.pop()
            self.s2.append(x)
        return self.s1.pop()

    def peek(self) -> int:
        # ** bug 3.2 - 忽略了s1为空的情况
        if not self.s1:
            # ** bug 4 - typo
            # return self.s2[0]
            return self.s2[-1]
        while len(self.s1) > 1:
            x = self.s1.pop()
            self.s2.append(x)
        return self.s1[-1]

    def empty(self) -> bool:
        # ** bug 1 - return type & logic
        # return self.s1 or self.s2
        return not self.s1 and not self.s2

# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
