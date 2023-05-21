package main

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	res := []string{}
	oneLine := []string{} //one line
	wordCount := 0
	for i := 0; i < len(words); i++ {
		cnt := len(words[i])
		if len(oneLine) > 0 {
			cnt++ //prefix space
		}
		if wordCount+cnt > maxWidth {
			s := produceOneline(oneLine, maxWidth)
			res = append(res, s)
			oneLine = []string{words[i]}
			wordCount = len(words[i])
		} else {
			oneLine = append(oneLine, words[i])
			wordCount += cnt
		}
	}
	//last line
	if len(oneLine) >= 1 {
		ss := strings.Join(oneLine, " ")
		spaces := getNSpace(maxWidth - len(ss))
		res = append(res, strings.Join([]string{ss, spaces}, ""))
	}
	return res
}

func produceOneline(oneline []string, maxWidth int) string {
	if len(oneline) == 0 {
		return ""
	} else if len(oneline) == 1 {
		//left positioned
		remaining := maxWidth - len(oneline[0])
		return oneline[0] + getNSpace(remaining)
	} else {
		remain := maxWidth
		for _, s := range oneline {
			remain -= len(s)
		}
		res := []string{}
		n := len(oneline) - 1
		for j := len(oneline) - 1; j > 0; j-- {
			res = append(res, oneline[j])
			sn := remain / n
			res = append(res, getNSpace(sn))
			n--
			remain -= sn
		}
		res = append(res, oneline[0])
		for i := 0; i < len(res)/2; i++ {
			res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
		}
		return strings.Join(res, "")
	}
}

func getNSpace(n int) string {
	spaces := make([]byte, n)
	for i := 0; i < n; i++ {
		spaces[i] = ' '
	}
	return string(spaces)
}
