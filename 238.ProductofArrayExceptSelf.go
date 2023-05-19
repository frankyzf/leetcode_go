package main

func productExceptSelf(nums []int) []int {
	total := 1
	for i := 0; i < len(nums); i++ {
		total *= nums[i]
		if total == 0 {
			break
		}
	}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		div := nums[i]
		if div != 0 {
			res = append(res, total/div)
		} else {
			subTotal := 1
			for j := 0; j < len(nums); j++ {
				if j != i {
					subTotal = subTotal * nums[j]
					if subTotal == 0 {
						break
					}
				}
			}
			res = append(res, subTotal)
		}
	}
	return res
}
