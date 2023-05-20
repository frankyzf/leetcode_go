package main

func isPalindrome(s string) bool {
	dst := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			dst = append(dst, s[i])
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			dst = append(dst, s[i]+32)
		} else if s[i] >= '0' && s[i] <= '9' {
			dst = append(dst, s[i])
		}
	}
	//finish convert string
	for i := 0; i < len(dst); i++ {
		j := len(dst) - i - 1
		if dst[i] != dst[j] {
			return false
		}
	}
	return true
}
