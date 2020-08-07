package main

import (
	"math"
	"sort"
)

func smallestDifference(array1, array2 []int) []int {
	// Write your code here.
	arr := make([]int, 0)
	// hashMap := make(map[int]int)

	diff := math.Abs(float64(array1[0]) - float64(array2[0]))
	arr = append(arr, array1[0], array2[0])
	for _, v1 := range array1 {
		for _, v2 := range array2 {
			// fmt.Println(diff, math.Abs(float64(v1)-float64(v2)), v1, v2)
			if math.Abs(float64(v1)-float64(v2)) < diff {
				arr[0], arr[1] = v1, v2
				diff = math.Abs(float64(v1) - float64(v2))
			}
		}
	}
	return arr
}

func smallestDifferencePointer(array1, array2 []int) []int {
	arr := make([]int, 0)
	sort.Ints(array1)
	sort.Ints(array2)
	pointer1, pointer2 := 0, 0
	diff := math.Abs(float64(array1[pointer1] - array2[pointer2]))
	arr = append(arr, array1[0], array2[0])
	for pointer1 < len(array1) && pointer2 < len(array2) {
		currDiff := math.Abs(float64(array1[pointer1] - array2[pointer2]))
		// fmt.Println(currDiff, array1[pointer1], array2[pointer2])
		if currDiff == 0 {
			return []int{array1[pointer1], array2[pointer2]}
		}
		if currDiff < diff {
			diff = currDiff
			arr[0], arr[1] = array1[pointer1], array2[pointer2]
		} else if array1[pointer1] < array2[pointer2] {
			pointer1++
		} else {
			pointer2++
		}
	}
	return arr
}
