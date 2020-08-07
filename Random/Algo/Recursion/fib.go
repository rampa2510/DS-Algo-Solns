package main

func fibRecursion(n int) int {
	if n == 2 {
		return 1
	}
	if n == 1 {
		return 0
	}

	return fibRecursion(n-1) + fibRecursion(n-2)
}
