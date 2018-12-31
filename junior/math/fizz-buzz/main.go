package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	cnt3, cnt5, rzt := 2, 2, []string{"1", "2"}
	if n <= 2 {
		return rzt
	}

	for i := 3; i <= n; i++ {
		cnt3++
		cnt5++
		flag := ""
		if cnt3 == 3 {
			flag = "Fizz"
			cnt3 = 0
		}
		if cnt5 == 5 {
			flag = "Buzz"
			cnt5 = 0
			if cnt3 == 0 {
				flag = "FizzBuzz"
			}
		}
		if cnt5 != 0 && cnt3 != 0 {
			flag = strconv.Itoa(i)
		}
		rzt = append(rzt, flag)
	}
	return rzt
}

func main() {
	fmt.Println(fizzBuzz(30))
}
