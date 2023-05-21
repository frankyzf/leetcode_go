package main

func maxArea(height []int) int {
	//outer loop from 0 to n -2
	//inner loop from n-1 to i+1
	if len(height) <= 1 {
		return 0
	}
	water := 0
	i, j := 0, len(height)-1
	for i < j {
		left := height[i]
		right := height[j]
		w := j - i
		h := left
		if h > right {
			h = right
		}
		a := h * w
		if a > water {
			water = a
		}
		if h == left {
			i++
		} else {
			j--
		}
	}
	return water
	//for each combination, calculate the water

}
