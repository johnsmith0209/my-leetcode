package main

import "fmt"

func romanToInt(s string) int {
	m, total, l := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}, 0, len(s)
	for i := 0; i < l; i++ {
		c := s[i]
		num := m[c]
		fmt.Printf("%s, %v, %d\n", string(c), c, num)
		if i < l-1 {
			next := m[s[i+1]]
			if next > num {
				total += next - num
				i++
			} else {
				total += num
			}
		} else {
			total += num
		}
	}
	return total
}

func main() {
	fmt.Println("IIIV", romanToInt("IIIV"))
}
