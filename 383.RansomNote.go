package main

func canConstruct(ransomNote string, magazine string) bool {
	stat := map[byte]int{}
	for _, b := range []byte(magazine) {
		stat[b] += 1
	}
	for _, b := range []byte(ransomNote) {
		n := stat[b]
		n -= 1
		if n < 0 {
			return false
		}
		stat[b] = n
	}
	return true
}
