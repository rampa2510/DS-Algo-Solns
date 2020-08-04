package main

func isPalindrome(k int) bool {

	var rev int = 0
	var r int = 0
	if k < 0 {
		return false
	}
	num := k
	for num > 0 {
		r = num % 10
		rev = rev*10 + r
		num = num / 10
	}
	return rev == k
}
