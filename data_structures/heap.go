package data_structures

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint64 |
		~float32 | ~float64
}

// fix: MaxHeap嵌入的MinHeap类型必须是结构体类型
// type MaxHeap[T Comparable] []T
type MaxHeap[T Comparable] struct {
	items []T
}

func (h MaxHeap[T]) Len() int {
	return len(h.items)
}
func (h MaxHeap[T]) Less(i, j int) bool {
	// ** i>j 最大堆
	return h.items[i] > h.items[j]
}
func (h MaxHeap[T]) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// cannot use h (variable of type *MaxHeap[int64]) as heap.Interface value in argument
// to heap.Push: *MaxHeap[int64] does not implement heap.Interface
// (wrong type for method Push)
// have Push(int64)
// want Push(any)
// func (h *MaxHeap[T]) Push(x T) {
func (h *MaxHeap[T]) Push(x any) {
	h.items = append(h.items, x.(T))
}

// func (h *MaxHeap[T]) Pop() T {
func (h *MaxHeap[T]) Pop() any {
	n := len(h.items)
	// old := h.items
	// x := old[n-1]
	// old[n-1] = *new(T)
	// h.items = old[:n-1]
	x := h.items[n-1]
	h.items = h.items[:n-1]
	return x
}

type MinHeap[T Comparable] struct {
	MaxHeap[T]
}

func (h MinHeap[T]) Less(i, j int) bool {
	// ** i<j 最小堆
	return h.items[i] < h.items[j]
}
