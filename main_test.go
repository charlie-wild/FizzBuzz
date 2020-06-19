package main

import "testing"

func TestItReturnsANumber(t *testing.T) {
	fizzBuzzResult := runFizzBuzz(1)

	if fizzBuzzResult != "1" {
		t.Errorf("Expected 1, got %s", fizzBuzzResult)
	}
}

func TestItReturnsFizzWhenPassed3(t *testing.T) {
	fizzBuzzResult := runFizzBuzz(3)

	if fizzBuzzResult != "Fizz" {
		t.Errorf("Expected Fizz, got %s", fizzBuzzResult)
	}
}

// // func itReturnsBuzzWhenPassed5(t *testing.T) {
// // 	fizzBuzzResult := runFizzBuzz()
// // }
