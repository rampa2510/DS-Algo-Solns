package main

func fourSum(array []int, target int) [][]int {
	arr := make([][]int, 0)
	hashMap := make(map[int][][]int)
	sum := 0
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			sum = array[i] + array[j]
			diff := target - sum
			if v, found := hashMap[diff]; found {
				for _, val := range v {
					newArr := append(val, array[i], array[j])
					arr = append(arr, newArr)
				}
				// continue
			}
		}
		for k := 0; k < i; k++ {
			sum = array[i] + array[k]
			hashMap[sum] = append(hashMap[sum], []int{array[k], array[i]})
		}
		// fmt.Println(hashMap)
	}
	return arr
	// for _,v1 := array
}
