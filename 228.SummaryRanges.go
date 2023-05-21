package main

import "fmt"

func summaryRanges(nums []int) []string {
	res := []string{}
	n := len(nums)
	if n == 0 {
		return []string{}
	}
	moveValue := nums[0]
	beginValue := nums[0]
	for i := 1; i < n; i++ {
		if nums[i]-moveValue == 1 {
			moveValue = nums[i] //move ahead
		} else { //not consecutive
			if beginValue == moveValue {
				res = append(res, fmt.Sprintf("%d", beginValue))
			} else {
				res = append(res, fmt.Sprintf("%d->%v", beginValue, moveValue))
			}
			moveValue = nums[i]
			beginValue = nums[i]
		}
	}
	if beginValue == moveValue {
		res = append(res, fmt.Sprintf("%d", beginValue))
	} else {
		res = append(res, fmt.Sprintf("%d->%v", beginValue, moveValue))
	}
	return res
}
