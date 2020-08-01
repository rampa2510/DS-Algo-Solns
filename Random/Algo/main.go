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
	a := []int{5, 1, 22, 25, 6, -1, 8, 10}
	s := []int{5, 1, 22, 22, 25, 6, -1, 8, 10}
	h := isValidSubsequence(a, s)
	fmt.Println(h)
	////
}
