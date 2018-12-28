package main

import (
	"fmt"
	"math"
	"strconv"
)

/*
https://blog.golang.org/strings

Implement atoi which converts a string to an integer.
The function first dicards as many whitespace characters as necessary until the first non-
whitespace character is found. Then, starting from this character, takes an optional initial
plus or minus sign followed by as many numerical digits as possible, and interprets them as
a numerical value.
The string can contain additinal characters after those that form the integral number,
which are ignored and have no effect on the behavior of this funciton.
If the first sequence of non-whitespace characters in str is not a valid integral number, or if
no such sequence exists because either str is empty of it contains only whitespace
charaters, no conversion is performed.
If no valid conversion could be performed, a zero value is returned.
*/

func isNum(c rune) bool {
	return c >= 48 && c <= 57
}

func isSign(c rune) bool {
	return c == 43 || c == 45
}

func isWhitespace(c rune) bool {
	return c == 32
}

func myAtoi(str string) int {
	rzt, sign, num, start := []rune{}, 1, 0, false
	min, max := -2147483648, 2147483647
	for _, c := range str {
		if isWhitespace(c) {
			if start {
				break
			} else {
				continue
			}
		}
		if !isNum(c) && !isSign(c) {
			break
		}
		if isSign(c) && start {
			break
		}
		start = true
		if isNum(c) {
			if len(rzt) == 0 && c == '0' {
				continue
			} else {
				rzt = append(rzt, c)
			}
		}
		if isSign(c) && c == '-' {
			sign = -1
		}
	}
	l := len(rzt)
	for k, c := range rzt {
		digits := int(math.Pow10(l - 1 - k))
		n, _ := strconv.Atoi(string(c))
		fmt.Println(n, digits)
		num += digits * n
	}
	fmt.Println(rzt, sign, num)
	num *= sign
	if num > max {
		return max
	} else if num < min {
		return min
	}
	return num
}

func main() {
	// fmt.Println(myAtoi("42"), 42)
	// fmt.Println(myAtoi("  -42"), -42)
	// fmt.Println(myAtoi("  -4 adf 2"), -4)
	// fmt.Println(myAtoi("asdf 24 4"), 0)
	fmt.Println(myAtoi("200000000000000000000"), 0)
}
