package main

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	leftSum := []int{0} //left, [start, i)
	leftLow := []int{0}
	rightSum := []int{0} //right [i, end]
	rightLow := []int{0}
	sum := 0
	prevLow := 100000
	for i := 1; i < n; i++ {
		node := gas[i-1] - cost[i-1]
		sum += node
		leftSum = append(leftSum, sum)
		if sum < prevLow {
			prevLow = sum
		} else {
			//do nothing
		}
		leftLow = append(leftLow, prevLow)

		rightSum = append(rightSum, 0) //init
		rightLow = append(rightLow, 0)
	}
	sum = 0
	prevLow = 100000
	for j := n - 1; j >= 0; j-- {
		node := gas[j] - cost[j]
		sum += node
		rightSum[j] = sum
		if prevLow+node > node {
			prevLow = node
		} else {
			prevLow = prevLow + node
		}
		rightLow[j] = prevLow
	}
	if sum < 0 {
		return -1
	}
	for i := 0; i < n; i++ {
		if rightLow[i] < 0 {
			continue
		}
		if rightSum[i] < leftLow[i] {
			continue
		}
		return i
	}

	return -1
}
