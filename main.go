package main

import "strconv"

func main() {

}

func runFizzBuzz(s int) string {
	switch {
	case s%3 == 0:
		return "Fizz"
	default:
		return strconv.Itoa(s)
	}

}
