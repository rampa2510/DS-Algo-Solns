package main

import "fmt"

func main() {
	a := [4]int{2, 8, 15, 18}
	b := [4]int{5, 9, 12, 17}
	c := merge(a, b)
	fmt.Println(c)
}
