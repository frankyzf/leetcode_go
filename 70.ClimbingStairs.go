package main

func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	buf := []int{1, 1}
	for i := 2; i <= n; i++ {
		cnt := buf[i-2] + buf[i-1]
		buf = append(buf, cnt)
	}
	return buf[n]
}
