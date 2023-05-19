package main

func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}
	cnt := map[int]int{}
	max := 0
	for i := 0; i < len(citations); i++ {
		cnt[citations[i]] += 1
		if max < citations[i] {
			max = citations[i]
		}
	}
	h := []int{}
	for i := 0; i < max+1; i++ {
		h = append(h, 0)
	}
	h[0] = len(citations)
	for i := 1; i < max+1; i++ {
		h[i] = h[i-1] - cnt[i-1]
		if h[i] < i {
			return i - 1
		}
	}
	return max
}
