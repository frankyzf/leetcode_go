package main

import (
	"strings"
)

func reverseWords(s string) string {
	ss := strings.Trim(s, " ")
	res := strings.Split(ss, " ")
	// fmt.Printf("ss:%v, len:%v\n", res, len(res))
	res2 := []string{}
	for i := 0; i < len(res); i++ {
		// fmt.Printf("word:%v\n", res[i])
		if res[i] != "" {
			res2 = append(res2, res[i])
		}
	}
	n := len(res2)
	for i := 0; i < int(n/2); i++ {
		res2[i], res2[n-1-i] = res2[n-1-i], res2[i]
	}
	ss = strings.Join(res2, " ")
	return ss
}
