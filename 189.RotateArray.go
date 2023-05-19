package main

func rotate(nums []int, k int) {
	tn := len(nums)
	if k > tn {
		k = k % tn
	}
	left := []int{}
	for i := 0; i < k; i++ {
		left = append(left, 0)
	}
	for i := 0; i < k; i++ {
		left[k-i-1] = nums[tn-i-1]
	}
	for i := 0; i < tn-k; i++ {
		nums[tn-i-1] = nums[tn-k-i-1]
	}

	for i := 0; i < k; i++ {
		nums[i] = left[i]
	}
	return
}
