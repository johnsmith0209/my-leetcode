package main

import "fmt"

func isValid(s string) bool {
	stack, pair := []rune{}, make(map[rune]rune, 3)
	pair[')'] = '('
	pair[']'] = '['
	pair['}'] = '{'
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
		} else {
			last := stack[len(stack)-1]
			should, ok := pair[v]
			if ok {
				if last == should {
					// fmt.Println("Before slice", stack)
					stack = stack[0 : len(stack)-1]
					// fmt.Println("After slice", stack)
				} else {
					// fmt.Println("Not match", should, v)
					return false
				}
			} else {
				stack = append(stack, v)
			}
		}
	}
	if len(stack) > 0 {
		// fmt.Println("not empty", stack)
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
