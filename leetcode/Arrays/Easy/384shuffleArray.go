package main

import (
	"math/rand"
	"time"
)

type solution struct {
	originalArr []int
	arr         []int
}

func constructor(array []int) solution {
	arr1 := make([]int, len(array))
	arr2 := make([]int, len(array))
	copy(arr1, array)
	copy(arr2, array)
	return solution{originalArr: arr1, arr: arr2}
}

func (s *solution) reset() []int {
	copy(s.arr, s.originalArr)
	return s.arr
}

func (s *solution) shuffle() []int {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(s.arr); i++ {
		j := rand.Intn(len(s.arr) - i)
		s.arr[j], s.arr[i] = s.arr[i], s.arr[j]
	}

	return s.arr
}
