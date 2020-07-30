package main

func merge(a, b [4]int) [8]int {
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
