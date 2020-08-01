package main

func moveElementToEnd(array []int, toMove int) []int {
	// Write your code here.
	arr := make([]int, 0)
	noOfInstance := 0
	for _, v := range array {
		if v != toMove {
			arr = append(arr, v)
		} else {
			noOfInstance++
		}
	}

	for i := 0; i < noOfInstance; i++ {
		arr = append(arr, toMove)
	}
	return arr
}

func moveElementToEndPointer(array []int, toMove int) []int {
	lPointer, rPointer := 0, len(array)-1
	for lPointer < rPointer {
		if array[lPointer] == toMove && array[rPointer] != toMove {
			array[lPointer], array[rPointer] = array[rPointer], array[lPointer]
		} else if array[lPointer] != toMove {
			lPointer++
		} else if array[rPointer] == toMove {
			rPointer--
		}
	}
	return array
}
