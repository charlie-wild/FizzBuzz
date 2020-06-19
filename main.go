package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(runFizzBuzz(i))
	}

}

func runFizzBuzz(s int) string {
	switch {
	case s%3 == 0 && s%5 == 0:
		return "FizzBuzz"
	case s%3 == 0:
		return "Fizz"
	case s%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(s)
	}

}
