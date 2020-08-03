package main

import (
	"fmt"
	"math"
)

func subArr(array []int) []int {
	min, max, start, stop := math.MinInt32, math.MaxInt32, -1, -1
	for i := 0; i < len(array); i++ {
		if i == 0 {
			if array[i] > array[i+1] {
				start = 0
				min = array[i]
				max = array[i]
			}
			continue
		}
		if i == len(array)-1 {
			if array[i] < array[i-1] {
				stop = i
				if array[i] < min {
					min = array[i]
				} else if array[i] > max {
					max = array[i]
				}
			}
			continue
		}

		if array[i] < array[i-1] || array[i] > array[i+1] {
			if start == -1 {
				start = i
				min = array[i]
				max = array[i]
			} else {
				stop = i
				if min > array[i] {
					min = array[i]
				} else if max < array[i] {
					max = array[i]
				}
			}
		}
	}
	fmt.Println(min, max, start, stop)
	if min == math.MinInt32 {
		return []int{-1, -1}
	}
	for i, ele := range array {
		if ele > min {
			start = i
			break
		}
	}

	for i := len(array) - 1; i >= 0; i-- {
		if array[i] < max {
			stop = i
			break
		}
	}

	return []int{start, stop}
}
