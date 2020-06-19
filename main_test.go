package main

import (
	"testing"
)

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

func TestItReturnsBuzzWhenPassed5(t *testing.T) {
	fizzBuzzResult := runFizzBuzz(5)

	if fizzBuzzResult != "Buzz" {
		t.Errorf("Expected Buzz, got %s", fizzBuzzResult)
	}
}

func TestItReturnsFizzBuzzWhenPassed15(t *testing.T) {
	fizzBuzzResult := runFizzBuzz(15)
	if fizzBuzzResult != "FizzBuzz" {
		t.Errorf("Expected FizzBuzz, got %s", fizzBuzzResult)
	}
}

func TestRangeOfNumbers(t *testing.T) {
	correctAnswerSlice := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz"}
	answerSlice := []string{}
	for i := 1; i <= len(correctAnswerSlice); i++ {
		answerSlice = append(answerSlice, runFizzBuzz(i))
	}

	for i := 0; i < len(correctAnswerSlice); i++ {
		if correctAnswerSlice[i] != answerSlice[i] {
			t.Errorf("Expected %s, got %s", correctAnswerSlice[i], answerSlice[i])
		}
	}
}
