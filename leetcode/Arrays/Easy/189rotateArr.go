package main

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
