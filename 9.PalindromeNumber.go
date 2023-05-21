package main

import "fmt"

func isPalindrome2(x int) bool {
	s := fmt.Sprintf("%d", x)
	ss := []byte(s)
	for i := 0; i < int(len(ss)/2); i++ {
		if ss[i] != ss[len(ss)-i-1] {
			return false
		}
	}
	return true
}
