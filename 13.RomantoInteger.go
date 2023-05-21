package main

var greekMM2 map[byte]int = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	buf := []byte(s)
	total := 0
	for i := 0; i < len(buf); i++ {
		num := greekMM2[buf[i]]
		if i < len(buf)-1 {
			nextNum := greekMM2[buf[i+1]]
			if nextNum > num {
				num = num * -1
			}
		}
		total += num
	}
	return total
}
