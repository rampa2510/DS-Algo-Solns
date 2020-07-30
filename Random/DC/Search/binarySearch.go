package main

import (
	"fmt"
	"sort"
)

func binSearch(a []int, low, up, target int) bool {
	if len(a) == 1 {
		return a[0] == target
	}
	if low > up {
		return false
	}
	mid := int((up + low) / 2)
	if a[mid] == target {
		return true
	}
	if a[mid] >= target {
		return binSearch(a, low, mid-1, target)
	}
	return binSearch(a, mid+1, up, target)

}

func main() {
	a := []int{3, 2, 1, 9, 4, 6, 7}
	sort.Ints(a)
	t := binSearch(a, 0, len(a), 8)
	fmt.Println(t)
}
