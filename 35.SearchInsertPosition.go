package main

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	start := 0
	end := len(nums) - 1
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	//start == end
	if nums[start] < target {
		return start + 1
	} else {
		return start
	}
}
