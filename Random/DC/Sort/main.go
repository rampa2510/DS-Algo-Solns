package main

import "fmt"

func main() {
	a := [4]int{2, 8, 15, 18}
	b := [4]int{5, 9, 12, 17}
	c := twoWayMerge(a, b)
	fmt.Println(c)
	d := []int{12, 11, 13, 5, 6, 7}
	mergeSort(d, 0, len(d)-1)
	fmt.Println(d)
}
