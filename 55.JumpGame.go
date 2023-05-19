package main

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	maxDst := nums[0]
	for i := 0; i <= maxDst && i < len(nums); i++ {
		d := i + nums[i]
		if maxDst < d {
			maxDst = d
		}
		if maxDst >= len(nums)-1 {
			return true
		}
	}
	return false
}
