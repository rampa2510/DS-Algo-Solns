package main

func isValidSubsequence(array []int, sequence []int) bool {
	arrIdx := 0
	seqIdx := 0
	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx++
		}
		arrIdx++
	}

	return seqIdx == len(sequence)
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func isValidate(array, sequence []int) bool {
	hashMap := make(map[int]int)
	for i, v := range sequence {
		hashMap[v] = i
	}
	c := 0
	prevIdx := -1
	for _, v := range array {
		if idx, found := hashMap[v]; found {
			if prevIdx > idx {
				return false
			}
			sequence = removeIndex(sequence, idx-c)
			prevIdx = idx
			c++
		}
	}
	if len(sequence) > 0 {
		return false
	}
	return true
}
