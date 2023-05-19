package main

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	MAX := 10000
	minSteps := []int{}
	for i := 0; i < len(nums); i++ {
		minSteps = append(minSteps, MAX)
	}
	minSteps[0] = 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if (j+nums[j]) >= i && minSteps[i] > minSteps[j]+1 {
				minSteps[i] = minSteps[j] + 1
			}
		}
	}
	return minSteps[len(nums)-1]
}
