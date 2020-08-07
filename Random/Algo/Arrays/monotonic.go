package main

func isMonotonic(array []int) bool {
	flag := 0
	if len(array) <= 1 {
		return true
	}
	prevElement := array[0]
	for i := 1; i < len(array); i++ {
		if prevElement != array[i] {
			if prevElement > array[i] {
				flag = 1
			}
		}
	}

	for i := 1; i < len(array); i++ {
		// fmt.Println(prevElement, array[i], flag)
		if flag == 1 {
			if prevElement < array[i] {
				return false
			}
		} else {
			if prevElement > array[i] {
				return false
			}
		}
		prevElement = array[i]
	}
	return true
}

func isMonotonicWfor(array []int) bool {
	if len(array) <= 2 {
		return true
	}
	flag := 0
	prevEle := array[0]
	for i := 1; i < len(array); i++ {
		if flag == 0 && prevEle != array[i] {
			if prevEle > array[i] {
				flag = 1
			} else {
				flag = 2
			}
			prevEle = array[i]
		} else if flag == 1 || flag == 2 {
			if flag == 1 && prevEle < array[i] {
				return false
			} else if flag == 2 && prevEle > array[i] {
				return false
			}
			prevEle = array[i]
		}
	}
	return true
}
