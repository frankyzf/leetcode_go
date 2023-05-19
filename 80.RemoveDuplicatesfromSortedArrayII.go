package main

func removeDuplicates2(nums []int) int {
	pos := 0
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[pos] == nums[i] {
			cnt++
			if cnt <= 2 {
				pos++
				nums[pos] = nums[i]
			}
		} else {
			pos++
			nums[pos] = nums[i]
			cnt = 1
		}
	}
	pos++
	return pos
}
