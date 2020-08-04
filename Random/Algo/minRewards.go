package main

func minRewards(scores []int) int {
	// Write your code here.
	rewards := make([]int, len(scores))
	fill(rewards, 1)
	for i := 1; i < len(rewards); i++ {
		if scores[i] > scores[i-1] {
			rewards[i] = rewards[i-1] + 1
			continue
		}
		for j := i - 1; j >= 0 && scores[j] > scores[j+1]; j-- {
			rewards[j] = max(rewards[j], rewards[j+1]+1)
		}
	}
	sum := 0
	for _, v := range rewards {
		sum += v
	}
	return sum
}

func fill(arr []int, j int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = j
	}
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
