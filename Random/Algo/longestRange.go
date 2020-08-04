package main

import (
	"fmt"
	"sort"
)

func longestRange(array []int) []int {
	// if len(array) == 1 {
	// 	return []int{array[0], array[0]}
	// }
	var best []int
	longestLen := 0
	hashMap := make(map[int]bool)

	for _, v := range array {
		hashMap[v] = false
	}

	for _, v := range array {
		if hashMap[v] {
			continue
		}

		currLen, left, right := 1, v-1, v+1
		hashMap[v] = true
		for hashMap[left] {
			currLen++
			hashMap[left] = true
			left--

		}

		for hashMap[right] {
			currLen++
			hashMap[right] = true
			right++

		}

		if currLen > longestLen {
			longestLen = currLen
			best = []int{left + 1, right - 1}
		}
	}
	return best

}

func sort2(array []int) []int {
	if len(array) == 1 {
		return []int{array[0], array[0]}
	}
	sort.Ints(array)
	newArr := make([]int, 0)
	hashMap := make(map[int]bool)
	for _, v := range array {
		if _, found := hashMap[v]; !found {
			newArr = append(newArr, v)
			hashMap[v] = true
		}
	}
	fmt.Println(newArr)
	var best []int
	longestLen, currLen := 0, 0
	for i := 1; i < len(newArr)-1; i++ {

		if newArr[i-1]+1 == newArr[i] {
			currLen++
			fmt.Println(newArr[i])
		} else {
			if currLen > longestLen {
				longestLen = currLen
				fmt.Println("a", currLen)
				best = []int{newArr[i-currLen-1], newArr[i-1]}
			}
			currLen = 0
			fmt.Println("l")
		}
	}
	return best
}
