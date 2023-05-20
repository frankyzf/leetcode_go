package main

func addBinary(a string, b string) string {
	abuf := []byte(a)
	bbuf := []byte(b)
	i, j := len(abuf), len(bbuf)
	if i == 0 {
		return b
	}
	if j == 0 {
		return a
	}
	res := []byte{}
	bb := byte('b')
	carry := 0
	for i, j = i-1, j-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		bb, carry = bitAdd(abuf[i], bbuf[j], carry)
		res = append([]byte{bb}, res...)
	}
	for ; i >= 0; i-- {
		bb, carry = bitAdd(abuf[i], '0', carry)
		res = append([]byte{bb}, res...)
	}
	for ; j >= 0; j-- {
		bb, carry = bitAdd('0', bbuf[j], carry)
		res = append([]byte{bb}, res...)
	}
	if carry == 1 {
		res = append([]byte{'1'}, res...)
	}
	return string(res)
}

func bitAdd(a byte, b byte, carry int) (byte, int) {
	if a == '1' {
		if b == '1' {
			if carry == 1 {
				return '1', 1
			} else {
				return '0', 1
			}
		} else {
			if carry == 1 {
				return '0', 1
			} else {
				return '1', 0
			}
		}
	} else {
		if b == '1' {
			if carry == 1 {
				return '0', 1
			} else {
				return '1', 0
			}
		} else {
			if carry == 1 {
				return '1', 0
			} else {
				return '0', 0
			}
		}
	}
}
