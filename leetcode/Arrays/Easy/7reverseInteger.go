package main

import (
	"math"
	"strconv"
)

func reverse(k int) int {

	strNum := []byte(strconv.Itoa(int(k)))

	if k < 0 {
		strNum = strNum[1:]
	}
	lenOfNum, end := len(strNum), len(strNum)-1

	for i := 0; i < lenOfNum/2; i, end = i+1, end-1 {
		strNum[i], strNum[end] = strNum[end], strNum[i]
	}
	num, _ := strconv.Atoi(string(strNum))
	if float64(num) > math.Pow(2, 31)-1 || float64(num) < -1*math.Pow(2, 31) {
		return 0
	}

	if k < 0 {
		return num * -1
	}
	return num
}

func reverseWithoutStr(k int) int {
	var isNeg bool = k < 0

	var rev int = 0
	var r int = 0
	if isNeg {
		k = k * -1
	}
	for k > 0 {
		r = k % 10
		rev = rev*10 + r
		k = k / 10
	}
	if isNeg {
		rev = rev * -1
	}

	if float64(rev) > math.Pow(2, 31)-1 || float64(rev) < -1*math.Pow(2, 31) {
		return 0
	}

	return rev
}
