package main

import (
	"fmt"
	"strings"
)

var sortArray []int = []int{1, 5, 10, 50, 100, 500, 1000}

var greekMM map[int]string = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	50:   "L",
	100:  "C",
	500:  "D",
	1000: "M",
}

func intToRoman(num int) string {
	res := []string{}
	for num > 0 {
		newPos := findPos(num)
		if newPos%2 == 1 {
			newPos--
		}
		newBase := sortArray[newPos]
		d := int(num / newBase)
		num = num - newBase*d
		if d < 4 {
			buf := ""
			for j := d; j > 0; j-- {
				buf += greekMM[sortArray[newPos]]
			}
			res = append(res, buf)
		} else if d == 4 {
			res = append(res, fmt.Sprintf("%v%v", greekMM[sortArray[newPos]], greekMM[sortArray[newPos+1]]))
		} else if d == 5 {
			res = append(res, fmt.Sprintf("%v", greekMM[sortArray[newPos+1]]))
		} else if d > 5 && d < 9 {
			buf := greekMM[sortArray[newPos+1]]
			d -= 5
			for j := d; j > 0; j-- {
				buf += greekMM[sortArray[newPos]]
			}
			res = append(res, buf)
		} else if d == 9 {
			res = append(res, fmt.Sprintf("%v%v", greekMM[sortArray[newPos]], greekMM[sortArray[newPos+2]]))
		}

	}
	return strings.Join(res, "")

}

func findPos(num int) int {
	for i, v := range sortArray {
		if v == num {
			return i
		} else if v > num {
			return i - 1 //first greater than num
		} else {
			//do nothing
		}
	}
	return len(sortArray) - 1
}
