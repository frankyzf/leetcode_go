package main

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	res := make([]int, len(ratings))
	for i, _ := range res {
		res[i] = 1
	}
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
		}
	}
	for j := len(ratings) - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] {
			if res[j] <= res[j+1] {
				res[j] = res[j+1] + 1
			}
		}
	}
	sum := 0
	for _, val := range res {
		sum += val
	}
	return sum
}
