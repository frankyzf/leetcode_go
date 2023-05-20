package main

var pair map[byte]byte = map[byte]byte{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	buf := []byte{}
	ss := []byte(s)
	for i := 0; i < len(ss); i++ {
		if b, ok := pair[ss[i]]; ok {
			if len(buf) > 0 && buf[len(buf)-1] == b {
				buf = buf[:len(buf)-1]
			} else {
				buf = append(buf, ss[i])
			}
		} else {
			buf = append(buf, ss[i])
		}
	}
	return len(buf) == 0
}
