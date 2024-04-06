package main

import "fmt"

// 【最短最难的面试题】深刻理解go语言切片的本质 #go #golang #go语言
// https://www.ixigua.com/7346719740179513907?logTag=d0815e41391c07965ab4

func main() {
	var arr [3]int
	brr := arr[1:2]
	brr = append(brr, 9)
	brr = append(brr, 9)
	fmt.Println(arr, len(arr), cap(arr))
	fmt.Println(brr, len(brr), cap(brr))
}
