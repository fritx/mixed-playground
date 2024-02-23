package main

type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint64 |
		~float32 | ~float64
}
type MaxHeap[T Comparable] []T

func (h MaxHeap[T]) Len() int {
	return len(h)
}
func (h MaxHeap[T]) Less(i, j int) bool {
	// ** i<j 最小堆
	// return h[i] < h[j]
	// ** i>j 最大堆
	return h[i] > h[j]
}
func (h MaxHeap[T]) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// cannot use h (variable of type *MaxHeap[int64]) as heap.Interface value in argument
// to heap.Push: *MaxHeap[int64] does not implement heap.Interface
// (wrong type for method Push)
// have Push(int64)
// want Push(any)
// func (h *MaxHeap[T]) Push(x T) {
func (h *MaxHeap[T]) Push(x any) {
	// *h = append(*h, x)
	*h = append(*h, x.(T))
}

// func (h *MaxHeap[T]) Pop() T {
func (h *MaxHeap[T]) Pop() any {
	n := len(*h)
	old := *h
	x := old[n-1]
	// old[n-1] = *new(T)
	*h = old[:n-1]
	return x
}
