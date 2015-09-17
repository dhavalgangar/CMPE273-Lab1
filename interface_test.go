package main

import "testing"

//Testing
//Interface

func TestInterface(t *testing.T){

	c := Circle{7}
	r := Rectangle{6,6}
	
	answer1 := c.calculatePerimeter()
	answer2 := r.calculatePerimeter()
	
	if answer1 != 43.982297150257104 {
		t.Errorf("Wrong answer %d", answer1)
	}
	
	if answer2 != 24 {
		t.Errorf("Wrong answer %d", answer2)
	}
}
