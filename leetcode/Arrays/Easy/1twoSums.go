package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) [2]int {
	hashMap := make(map[int][2]int)
	lengthOfNums := len(nums) - 1
	for i := 0; i <= lengthOfNums; i++ {
		if _, found := hashMap[nums[i]]; found {
			if nums[i]+nums[i] == target {
				return [2]int{hashMap[nums[i]][1], i}
			}
		}
		hashMap[nums[i]] = [2]int{target - nums[i], i}
	}
	var indice [2]int

	for k, v := range hashMap {
		delete(hashMap, k)
		requiredNum := v[0]
		if requiredElement, isElementPresent := hashMap[requiredNum]; isElementPresent {
			indice[0] = v[1]
			indice[1] = requiredElement[1]
			break
		}
	}

	return indice
}

func twoSumSort(arr []int, target int) {
	sort.Ints(arr)
	lPointer := 0
	rPointer := len(arr) - 1
	for true {
		sum := arr[lPointer] + arr[rPointer]
		if sum == target {
			fmt.Println(arr[lPointer], arr[rPointer])
			break
		}
		if rPointer <= lPointer {
			// return []int{}
		}
		if sum > target {
			rPointer--
		} else {
			lPointer++
		}
	}
}
