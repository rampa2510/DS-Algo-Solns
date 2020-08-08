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
	// a := [][]int{[]int{1, 2, 3, 4}, []int{12, 13, 14, 5}, []int{11, 16, 15, 6}, []int{10, 9, 8, 7}}
	// b := spiralTraverse(a)
	// fmt.Println(b)
	////

	//// Longest Peak
	// a := []int{1, 2, 3, 4, 5, 1}
	// b := longestPeak(a)
	// fmt.Println(b)
	////

	//// FOur sum
	// a := []int{7, 6, 4, -1, 1, 2}
	// b := sum(a, 16)
	// fmt.Println(b)
	////

	//// Subarr sort
	// a := []int{1, 2, 4, 7, 10, 11, 7, 12, 6, 7, 16, 18, 19}
	// b := subArray(a)
	// fmt.Println(b)
	////

	//// Longest Range
	// a := []int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6}
	// b := longestRange(a)
	// fmt.Println(b)
	////

	//// Min rewards
	// a := []int{2, 20, 13, 12, 11, 8, 4, 3, 1, 5, 6, 7, 9, 0}
	// b := minRewards(a)
	// fmt.Println(b)
	////

	/// Is validate hashmap
	// a := []int{5, 1, 22, 25, 6, -1, 8, 10}
	// b := []int{5, 1, 22, 25, 6, -1, 8, 10, 12}
	// c := isValidate(a, b)
	// fmt.Println(c)
	///

	////zigzag
	a := [][]int{
		{1, 3, 4, 10},
		{2, 5, 9, 11},
		{6, 8, 12, 15},
		{7, 13, 14, 16},
	}
	b := zig(a)
	fmt.Println(b)

}
