package main

import (
	"fmt"
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

func main() {
	reverseNum := reverse(123)
	fmt.Println(reverseNum)
}
