package main

import "fmt"

type strPos struct {
	StrIdx, Count int
}

func firstUniqChar(s string) int {
	count := make(map[rune]strPos, len(s))
	for k, v := range s {
		info, ok := count[v]
		if ok {
			// rm list index
			info.Count++
			count[v] = info
			fmt.Println("increase", count[v])
		} else {
			count[v] = strPos{
				StrIdx: k,
				Count:  1,
			}
			fmt.Println("create node, append to list", count[v])
		}
	}
	idx := len(s) + 1
	for _, v := range count {
		if v.Count == 1 && v.StrIdx < idx {
			idx = v.StrIdx
		}
	}
	return idx
}

func main() {
	// a := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(a[2+1:])
	// fmt.Println(append(a[:2], a[3:]...))

	s := "footballfetter"
	fmt.Println(firstUniqChar(s))

	fmt.Println('A' - 'b')
}
