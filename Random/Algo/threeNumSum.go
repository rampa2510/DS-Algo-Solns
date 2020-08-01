package main

import (
	"sort"
)

func threeNumberSum(array []int, target int) [][]int {
	arr := make([][]int, 0)
	sort.Ints(array)
	for _, v1 := range array {
		for _, v2 := range array {
			for _, v3 := range array {
				if v1+v2+v3 == target && v3 > v2 && v2 > v1 {
					arr = append(arr, []int{v1, v2, v3})
				}
			}
		}
	}

	return arr

}

func threeNumberSumHash(array []int, target int) [][]int {
	hashMap := make(map[int]int)
	arr := make([][]int, 0)
	sort.Ints(array)
	for _, v := range array {
		hashMap[v] = v
	}
	sum := 0
	for _, v1 := range array {
		for _, v2 := range array {
			if v1 < v2 {
				sum = v1 + v2
				if v3, found := hashMap[target-sum]; found {
					if v3 > v2 {
						arr = append(arr, []int{v1, v2, v3})
					}
				}
			}
		}
	}
	return arr
}

func threeNumberSumPointer(array []int, target int) [][]int {
	arr := make([][]int, 0)
	sort.Ints(array)
	lPointer := 0
	rPointer := len(array) - 1
	diff := 0
	sum := 0
	for i, v1 := range array {
		diff = target - v1
		for lPointer, rPointer = i, len(array)-1; lPointer < rPointer; {
			// fmt.Println(lPointer, rPointer, diff, sum)
			sum = array[lPointer] + array[rPointer]
			if sum == diff && array[lPointer] < array[rPointer] && v1 < array[lPointer] {
				arr = append(arr, []int{v1, array[lPointer], array[rPointer]})
			}
			if sum > diff {
				rPointer--
			} else {
				lPointer++
			}

		}
	}
	return arr

}
