package main

func twoWayMerge(a, b [4]int) [8]int {
	i, j, k := 0, 0, 0
	c := [8]int{}

	for i < len(a) && j < len(b) {
		if a[i] > b[j] {
			c[k] = b[j]
			j, k = j+1, k+1
		} else {
			c[k] = a[i]
			i, k = i+1, k+1
		}
	}
	for i < len(a) {
		c[k] = a[i]
		i, k = i+1, k+1
	}
	for j < len(b) {
		c[k] = b[j]
		j, k = j+1, k+1
	}
	return c
}

func mergeSort(a []int, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeSort(a, low, mid)
	mergeSort(a, mid+1, high)
	merge(a, low, mid, high)
}

func merge(a []int, l, m, h int) {

	i, j, k := l, m+1, 0
	temp := [10]int{}

	for i <= m && j <= h {
		if a[i] > a[j] {
			temp[k] = a[j]
			j, k = j+1, k+1
		} else {
			temp[k] = a[i]
			i, k = i+1, k+1
		}
	}
	// fmt.Println(temp, l, h, m)
	for i <= m {
		temp[k] = a[i]
		i, k = i+1, k+1
	}

	for j <= h {
		temp[k] = a[j]
		j, k = j+1, k+1
	}

	for i, k = l, 0; i <= h; i, k = i+1, k+1 {
		a[i] = temp[k]
	}
	// fmt.Println(a)
}
