package main

import "strings"

func lengthOfLastWord(s string) int {
	s2 := []byte(strings.Trim(s, " "))
	if len(s2) == 0 {
		return 0
	}
	j := len(s2) - 1
	for ; j >= 0; j-- {
		if s2[j] == ' ' {
			j++
			break
		} else {
			//keep looping
		}
	}
	if j < 0 {
		j = 0
	}
	return len(s2) - j

}
