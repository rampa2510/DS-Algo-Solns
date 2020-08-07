package main

import "fmt"

func main() {
	b := SpecialArray{SpecialArray{1, 2}}
	a := productSum(b)
	fmt.Println(a)
}
