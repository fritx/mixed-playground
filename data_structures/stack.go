package data_structures

type Stack[T any] interface {
	Push(x T)
	Pop() T
	Top() T
	Empty() bool
}
