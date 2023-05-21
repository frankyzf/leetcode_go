package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	end := 0
	bLoop := true
	for bLoop {
		var digit byte
		for i, str := range strs {
			if end >= len(str) {
				bLoop = false
				break
			} else { //no hit word end
				if i == 0 {
					digit = str[end]
				} else {
					if digit != str[end] {
						bLoop = false
						break
					} else {
						//keep loop
					}
				}
			}
		}
		end++
	}
	if end == 1 {
		return ""
	} else {
		return strs[0][0 : end-1]
	}
}
