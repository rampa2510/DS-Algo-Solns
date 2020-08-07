package main

import (
	"math"
)

// func subArr(array []int) []int {
// 	min, maxNum, start, stop := math.MinInt32, math.MaxNumInt32, -1, -1
// 	for i := 0; i < len(array); i++ {
// 		if i == 0 {
// 			if array[i] > array[i+1] {
// 				start = 0
// 				min = array[i]
// 				maxNum = array[i]
// 			}
// 			continue
// 		}
// 		if i == len(array)-1 {
// 			if array[i] < array[i-1] {
// 				stop = i
// 				if array[i] < min {
// 					min = array[i]
// 				} else if array[i] > maxNum {
// 					maxNum = array[i]
// 				}
// 			}
// 			continue
// 		}

// 		if array[i] < array[i-1] || array[i] > array[i+1] {
// 			if start == -1 {
// 				start = i
// 				min = array[i]
// 				maxNum = array[i]
// 			} else {
// 				stop = i
// 				if min > array[i] {
// 					min = array[i]
// 				} else if maxNum < array[i] {
// 					maxNum = array[i]
// 				}
// 			}
// 		}
// 	}
// 	fmt.Println(min, maxNum, start, stop)
// 	if min == math.MinInt32 {
// 		return []int{-1, -1}
// 	}
// 	for i, ele := range array {
// 		if ele > min {
// 			start = i
// 			break
// 		}
// 	}

// 	for i := len(array) - 1; i >= 0; i-- {
// 		if array[i] < maxNum {
// 			stop = i
// 			break
// 		}
// 	}

// 	return []int{start, stop}
//}

func isOutOfOrder(i, num int, arr []int) bool {
	if i == 0 {
		return num > arr[1]
	}
	if i == len(arr)-1 {
		return arr[i-1] > num
	}

	return arr[i] > arr[i+1] || arr[i] < arr[i-1]
}

func min(i, j int) int {
	if i > j {
		return j
	}

	return i
}

// func max(i, j int) int {
// 	if i > j {
// 		return i
// 	}

// 	return j
// }

func subArray(array []int) []int {
	minOutOfOrder, maxOutOfOrder := math.MaxInt32, math.MinInt16
	for i, num := range array {
		if isOutOfOrder(i, num, array) {
			minOutOfOrder = min(minOutOfOrder, num)
			maxOutOfOrder = max(maxOutOfOrder, num) // definition of max minRewards.go file
		}
	}
	if minOutOfOrder == math.MaxInt32 {
		return []int{-1, -1}
	}
	left := 0
	for minOutOfOrder >= array[left] {
		left++
	}
	right := len(array) - 1
	for maxOutOfOrder <= array[right] {

		right--
	}

	return []int{left, right}
}
