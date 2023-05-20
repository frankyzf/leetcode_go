package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	sbuf := []byte(s)
	tbuf := []byte(t)
	for round := 0; round <= len(t)-len(s); round++ {
		if sbuf[0] == tbuf[round] {
			if roundCompare(sbuf, tbuf, round) {
				return true
			}
		} else {
			//continue, do nothing
		}
	}
	return false
}

func roundCompare(sbuf, tbuf []byte, ti int) bool {
	si := 0
	for ; si < len(sbuf) && ti < len(tbuf); ti++ {
		if sbuf[si] == tbuf[ti] {
			si++
		}
	}
	return si == len(sbuf)
}
