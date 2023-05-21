package main

func trap(height []int) int {
	//1. setup two array, leftHigh and rightHigh to track the unit left and right high
	n := len(height)
	if n == 0 {
		return 0
	}
	leftHigh := make([]int, n)
	rightHigh := make([]int, n)
	highest := 0
	for i, h := range height {
		if h > highest {
			highest = h
		}
		leftHigh[i] = highest
	}
	highest = 0
	for j := n - 1; j >= 0; j-- {
		h := height[j]
		if h > highest {
			highest = h
		}
		rightHigh[j] = highest
	}
	//2.confirm the unit by get lower of left/right high
	//3. the trapped water is the confirmed unit - volume of the pilar
	sum := 0
	for i := 0; i < n; i++ {
		rh := leftHigh[i]
		if rh > rightHigh[i] {
			rh = rightHigh[i]
		}
		sum += (rh - height[i])
	}
	return sum
}
