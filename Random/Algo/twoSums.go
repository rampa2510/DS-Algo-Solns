package main

import (
	"fmt"
	"sort"
)

// Question - Given a array find 2 numbers whithin the array that add up to a target

func twoNumberSumHashMap(array []int, target int) []int {
	// Write your code here.
	hashMap := make(map[int]int)
	for _, v := range array {
		if _, found := hashMap[v]; found && v+v == target {
			return []int{v, v}
		}
		if element, found := hashMap[target-v]; found {
			return []int{element, v}
		}
		hashMap[v] = v
	}

	return []int{}
}

func twoSumSort(arr []int, target int) []int {
	sort.Ints(arr)
	lPointer := 0
	rPointer := len(arr) - 1
	for true {
		sum := arr[lPointer] + arr[rPointer]
		if sum == target {
			return []int{arr[lPointer], arr[rPointer]}
		}
		if rPointer <= lPointer {
			return []int{}
		}
		if sum > target {
			rPointer--
		} else {
			lPointer++
		}
	}
	return []int{}
}

func twoSumBrute(array []int, target int) []int {
	// Write your code here.
	for i1, num1 := range array {
		for i2, num2 := range array {
			if num1+num2 == target && i1 != i2 {
				return []int{num1, num2}
			}
		}
	}
	return []int{}
}

func main() {
	a := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	hashMapSoln := twoNumberSumHashMap(a, target)
	sortSoln := twoSumSort(a, target)
	bruteSoln := twoSumBrute(a, target)
	fmt.Println(hashMapSoln, sortSoln, bruteSoln)
}
