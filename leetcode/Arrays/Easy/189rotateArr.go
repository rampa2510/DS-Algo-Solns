package main

import "fmt"

func rotate(nums []int, k int) {
	i := 0
	lengthOfNums := len(nums) - 1
	for i < k {
		temp := nums[lengthOfNums]
		for j := lengthOfNums; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = temp
		i++
	}

}

func main() {
	prob := make([]int, 0)
	prob = append(prob, 1, 2, 3, 4, 5, 6, 7)
	rotate(prob, 3)
	fmt.Println(prob)
}
