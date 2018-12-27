package main

import "fmt"

func isUpperCase(c byte) bool {
	return c >= 65 && c <= 90
}

func isLowerCase(c byte) bool {
	return c >= 97 && c <= 122
}

func isNum(c byte) bool {
	return c >= 48 && c <= 57
}

func isPalindrome(s string) bool {
	idx1, idx2 := 0, len(s)-1
	for idx1 < idx2 {
		char1, char2 := s[idx1], s[idx2]
		if !isUpperCase(char1) && !isLowerCase(char1) && !isNum(char1) {
			fmt.Println(char1, "1 not a-zA-Z0-9 continue", idx1, idx1+1)
			idx1++
			continue
		}
		if !isUpperCase(char2) && !isLowerCase(char2) && !isNum(char2) {
			fmt.Println(char2, "2 not a-zA-Z0-9 continue", idx2, idx2-1)
			idx2--
			continue
		}
		if isUpperCase(char1) {
			fmt.Println(char1, "1 is upper case, to lower", char1+32)
			char1 += 32
		}
		if isUpperCase(char2) {
			fmt.Println(char2, "2 is upper case, to lower", char2+32)
			char2 += 32
		}
		if char1 != char2 {
			fmt.Println(char1, char2, "not equal quit")
			return false
		} else {
			fmt.Println(char1, char2, "equal continue")
			idx1++
			idx2--
		}
	}
	return true
}

func main() {
	// fmt.Println("!!!!result=================== A man, a plan, a canal: Panama ", isPalindrome("A man, a plan, a canal: Panama"), true)
	// fmt.Println("!!!!result=================== Aa", isPalindrome("Aa"), true)
	// fmt.Println("!!!!result=================== A a", isPalindrome("A a"), true)
	// fmt.Println("!!!!result=================== A b", isPalindrome("A b"), false)
	// fmt.Println("!!!!result=================== A b a", isPalindrome("A b a"), true)
	fmt.Println("!!!!result=================== 0P", isPalindrome("0P"), false)
	fmt.Println("!!!!result=================== 0P", isPalindrome(" P"), true)
}
