package main

func longestPeak(array []int) int {
	// Write your code here.
	// arr := make([][]int, 0)
	longestPeak := 0
	for i := 1; i < len(array)-1; i++ {
		if array[i] > array[i-1] && array[i+1] < array[i] {
			leftIdx, rightIdx := i-2, i+2
			for leftIdx >= 0 && array[leftIdx] < array[leftIdx+1] {
				leftIdx--
			}
			for rightIdx < len(array) && array[rightIdx] < array[rightIdx-1] {
				rightIdx++
			}
			size := rightIdx - leftIdx - 1
			if longestPeak < size {
				longestPeak = size
			}
		}
	}

	// fmt.Println(peaks)
	return longestPeak
}
