package main

import "fmt"

func isAnagram(s string, t string) bool {
	m := map[rune]int{}
	for _, v := range s {
		_, ok := m[v]
		if ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for _, v := range t {
		_, ok := m[v]
		if ok {
			m[v]--
		} else {
			m[v] = -1
		}
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s, t := "", ""
	fmt.Println(isAnagram(s, t))
}
