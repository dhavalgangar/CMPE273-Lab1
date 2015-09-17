package main

import "testing"

//Testing
//Fibonacci series

func TestFibonnaci(t *testing.T) {
	answer := fibbo(6.0)
	if answer != 8.0 {
		t.Errorf("Wrong Fibonacci series answer %d", answer)
	}
}
