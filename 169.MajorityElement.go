package main

func majorityElement(nums []int) int {
	maxn, _ := maxNum(nums, 0, len(nums)-1)
	return maxn
}

func maxNum(nums []int, start, end int) (int, int) {
	if start == end {
		return nums[start], 1
	}
	//divide 2 parts
	n := int((end - start) / 2)
	left, leftn := maxNum(nums, start, start+n)
	right, rightn := maxNum(nums, start+n+1, end)
	for i := start + n + 1; i <= end; i++ {
		if nums[i] == left {
			leftn++
		}
	}
	for i := start + n; i >= start; i-- {
		if nums[i] == right {
			rightn++
		}
	}
	if leftn >= rightn {
		return left, leftn
	}
	return right, rightn
}
