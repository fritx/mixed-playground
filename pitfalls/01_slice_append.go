package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("-- ok")
	{
		a := []int{1, 2, 3}
		b := append(a, 8)
		c := append(a, 9)
		fmt.Printf("a=%v\n", a)
		fmt.Printf("b=%v\n", b)
		fmt.Printf("c=%v\n", c)
		fmt.Println("")
	}
	fmt.Println("-- buggy")
	{
		a := []int{}
		for i := 0; i < 5; i++ {
			fmt.Printf("a=%v\n", a)
			b := append(a, 11)
			c := append(a, 22)
			fmt.Printf("b=%v, c=%v\n", b, c)
			// ...
			// b=[0 1 11], c=[0 1 22]
			// 💥 b=[0 1 2 22], c=[0 1 2 22]
			// b=[0 1 2 3 11], c=[0 1 2 3 22]
			a = append(a, i)
		}
		fmt.Println("")
	}
	fmt.Println("-- buggy 2")
	{
		a := []int{}
		bb := [][]int{}
		cc := [][]int{}
		for i := 0; i < 5; i++ {
			fmt.Printf("a=%v\n", a)
			b := append(a, 11)
			c := append(a, 22)
			bb = append(bb, b)
			cc = append(cc, c)
			fmt.Printf("bb=%v, cc=%v\n", bb, cc)
			// ...
			// bb=[[11] [0 11]], cc=[[22] [0 22]]
			// bb=[[11] [0 11] [0 1 11] [0 1 2 22]], cc=[[22] [0 22] [0 1 22] [0 1 2 22]]
			// 💥 bb=[[11] [0 11] [0 1 11] [0 1 2 3] [0 1 2 3 11]],
			//   cc=[[22] [0 22] [0 1 22] [0 1 2 3] [0 1 2 3 22]]
			a = append(a, i)
		}
		fmt.Println("")
	}
	fmt.Println("-- inspect")
	{
		a := []int{}
		for i := 0; i < 5; i++ {
			fmt.Printf("a=%v, cap(a)=%v\n", a, cap(a))
			b := append(a, 11)
			c := append(a, 22)
			fmt.Printf("b=%v, c=%v\n", b, c)
			a = append(a, i)
		}
		fmt.Println("")
	}
	fmt.Println("-- fix via copy()")
	{
		a := []int{}
		for i := 0; i < 5; i++ {
			fmt.Printf("a=%v\n", a)
			aa := make([]int, len(a)) // cap=len while make()
			copy(aa, a)
			b := append(aa, 11)
			c := append(aa, 22)
			fmt.Printf("b=%v, c=%v\n", b, c)
			a = append(a, i)
		}
		fmt.Println("")
	}
	fmt.Println("-- fix via spreading")
	{
		a := []int{}
		for i := 0; i < 5; i++ {
			fmt.Printf("a=%v\n", a)
			ab := append([]int{}, a...)
			b := append(ab, 11)
			ac := append([]int{}, a...)
			c := append(ac, 22)
			fmt.Printf("b=%v, c=%v\n", b, c)
			a = append(a, i)
		}
		fmt.Println("")
	}
	// requires Go >=1.22
	fmt.Println("-- fix via slices.Concat()")
	{
		a := []int{}
		for i := 0; i < 5; i++ {
			fmt.Printf("a=%v\n", a)
			ab := slices.Concat[[]int]([]int{}, a)
			b := append(ab, 11)
			ac := slices.Concat[[]int]([]int{}, a)
			c := append(ac, 22)
			fmt.Printf("b=%v, c=%v\n", b, c)
			a = append(a, i)
		}
		fmt.Println("")
	}
	fmt.Println("-- fix via slicing with specifying cap")
	{
		a := []int{}
		for i := 0; i < 5; i++ {
			fmt.Printf("a=%v\n", a)
			b := append(a[:len(a):len(a)], 11)
			c := append(a[:len(a):len(a)], 22)
			fmt.Printf("b=%v, c=%v\n", b, c)
			a = append(a, i)
		}
		fmt.Println("")
	}
	fmt.Println("-- learn slices.Grow")
	{
		a := make([]int, 3, 7)
		b := Grow(a, 13)
		fmt.Printf("a=%v, len(a)=%v, cap(a)=%v\n", a, len(a), cap(a))
		fmt.Printf("b=%v, len(b)=%v, cap(b)=%v\n", b, len(b), cap(b))
	}
}

// ~go/slices/slices.go
// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation. If n is negative or too large to
// allocate the memory, Grow panics.
func Grow[S ~[]E, E any](s S, n int) S {
	if n < 0 {
		panic("cannot be negative")
	}
	if n -= cap(s) - len(s); n > 0 {
		s = append(
			s[:cap(s)],
			make([]E, n)..., // <-- expanding cap
		)[:len(s)] // <-- cutting len
	}
	return s
}

// Concat returns a new slice concatenating the passed in slices.
func Concat[S ~[]E, E any](slices ...S) S {
	size := 0
	for _, s := range slices {
		size += len(s)
		if size < 0 {
			panic("len out of range")
		}
	}
	newslice := Grow[S](nil, size)
	for _, s := range slices {
		newslice = append(newslice, s...)
	}
	return newslice
}
