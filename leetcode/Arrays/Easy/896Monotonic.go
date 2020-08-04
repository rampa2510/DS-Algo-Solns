package main

func isMonotonic(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}
	flag := 0

	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		if flag == 0 && prev != arr[i] {
			if prev > arr[i] {
				flag = 2
			} else {
				flag = 1
			}
		} else if flag == 1 {
			if prev > arr[i] {
				return false
			}
		} else if flag == 2 {
			if prev < arr[i] {
				return false
			}
		}
		prev = arr[i]
	}
	return true
}
