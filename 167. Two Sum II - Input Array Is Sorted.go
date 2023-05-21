package main

func twoSum2(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		sum := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			if sum+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{}
}
