package main

import "testing"

func itReturnsANumber(t *testing.T) {
	fizzBuzzResult := runFizzBuzz()

	if fizzBuzzResult != 1 {
		t.Errorf("Expected 1, got %d", fizzBuzzResult)
	}
}
