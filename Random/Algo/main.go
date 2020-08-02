package main

import "fmt"

func main() {
	/// Two number sum
	// a := []int{3, 5, -4, 8, 11, 1, -1, 6}
	// target := 10
	// hashMapSoln := twoNumberSumHashMap(a, target)
	// sortSoln := twoSumSort(a, target)
	// bruteSoln := twoSumBrute(a, target)
	// fmt.Println(hashMapSoln, sortSoln, bruteSoln)
	/////

	//// validate sequence
	// a := []int{5, 1, 22, 25, 6, -1, 8, 10}
	// s := []int{5, 1, 22, 22, 25, 6, -1, 8, 10}
	// h := isValidSubsequence(a, s)
	// fmt.Println(h)
	////

	/// three sum
	// a := []int{12, 3, 1, 2, -6, 5, -8, 6}
	// b := threeNumberSumPointer(a, 0)
	// fmt.Println(b)
	///

	//// Smallest Difference
	// a := []int{240, 124, 86, 111, 2, 84, 954, 27, 89}
	// b := []int{1, 3, 954, 19, 8}
	// c := smallestDifferencePointer(a, b)
	// fmt.Println(c)
	////

	/// Monotonic
	// a := []int{-1, -1, -2, -3, -4, -5, -5, -5, -6, -7, -8, -8, -9, -10, -11}
	// b := isMonotonicWfor(a)
	// fmt.Println(b)
	///

	//// Spiral Traverse
	a := [][]int{[]int{1, 2, 3, 4}, []int{12, 13, 14, 5}, []int{11, 16, 15, 6}, []int{10, 9, 8, 7}}
	b := spiralTraverse(a)
	fmt.Println(b)
	////
}
