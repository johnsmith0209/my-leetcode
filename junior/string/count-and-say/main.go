package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	rzt, s := "", ""
	for i := 0; i < n; i++ {
		if len(rzt) == 0 {
			rzt = "1"
			continue
		}
		num, idx, cnt := rzt[0], 0, 0
		for idx < len(rzt) {
			if num == rzt[idx] {
				cnt++
			} else {
				s += strconv.Itoa(cnt) + string(num)
				cnt, num = 1, rzt[idx]
			}
			idx++
		}
		if cnt > 0 {
			s += strconv.Itoa(cnt) + string(num)
		}
		rzt, s = s, ""
	}
	return rzt
}

func main() {
	fmt.Println(countAndSay(30))
}
