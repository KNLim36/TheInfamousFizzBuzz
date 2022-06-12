package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		result := ""
		if i%3 == 0 {
			result += "Fizz"
		}
		if i%5 == 0 {
			result += "Buzz"
		}
		if result == "" {
			result = strconv.Itoa(i)
		}
		fmt.Println(result)
	}
}
