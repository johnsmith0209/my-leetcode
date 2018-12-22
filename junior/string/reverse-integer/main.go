package main

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {
	const intMax = int(^uint32(0) >> 1)
	const intMin = ^intMax
	if x > intMax || x < intMin {
		return 0
	}
	negative := x < 0
	if negative {
		x *= -1
	}
	str := strconv.Itoa(x)
	l := len(str)
	strRune := make([]rune, l)
	for k, v := range str {
		strRune[k] = v
	}
	for i := 0; i < l/2; i++ {
		strRune[i], strRune[l-1-i] = strRune[l-1-i], strRune[i]
	}
	rzt, _ := strconv.Atoi(string(strRune))
	if negative {
		rzt *= -1
	}
	return rzt
}

func main() {
	fmt.Println(reverse(134))
}
