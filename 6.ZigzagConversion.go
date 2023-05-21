package main

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	n := len(s)
	res := make([]byte, n)
	buf := []byte(s)
	round := numRows*2 - 2
	cnt := 0
	for i := 0; i < numRows; i++ { //line by line
		for j := i; j < n; j += round {
			res[cnt] = buf[j]
			cnt++
			if i != 0 && i != numRows-1 {
				//etra line
				extraLine := round - 2*i
				if j+extraLine < n {
					res[cnt] = buf[j+extraLine]
					cnt++
				}
			}
		}
	}
	return string(res)
}
