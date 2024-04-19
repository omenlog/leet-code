package main

import "strings"

func IntToRoman(num int) string {
	res := ""

	x := num
	res += strings.Repeat("M", x/1000)

	x %= 1000
	res += strings.Repeat("CM", x/900)

	x %= 900
	res += strings.Repeat("D", x/500)

	x %= 500
	res += strings.Repeat("CD", x/400)

	x %= 400
	res += strings.Repeat("C", x/100)

	x %= 100
	res += strings.Repeat("XC", x/90)

	x %= 90
	res += strings.Repeat("L", x/50)

	x %= 50
	res += strings.Repeat("XL", x/40)

	x %= 40
	res += strings.Repeat("X", x/10)

	x %= 10
	res += strings.Repeat("IX", x/9)

	x %= 9
	res += strings.Repeat("V", x/5)

	x %= 5
	res += strings.Repeat("IV", x/4)

	x %= 4
	res += strings.Repeat("I", x/1)

	return res
}
