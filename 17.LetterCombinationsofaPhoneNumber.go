package main

var mm map[byte]string = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	container := []string{}
	value := []byte(mm[digits[0]])
	for j := 0; j < len(value); j++ {
		container = append(container, string(value[j]))
	}
	for i := 1; i < len(digits); i++ {
		container = genOneRound(container, digits[i])
	}
	return container

}

func genOneRound(buf []string, digit byte) []string {
	res := []string{}
	for i := 0; i < len(buf); i++ {
		item := []byte(buf[i])
		value := []byte(mm[digit])
		for j := 0; j < len(value); j++ {
			newitem := append(item, value[j])
			res = append(res, string(newitem))
		}
	}
	return res
}
