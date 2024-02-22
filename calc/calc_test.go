package main

import (
	"testing"
)

func TestCheckSuccess(t *testing.T) {
	var inputString = "1*(6*8*6)*(18/2/3/3*9)+4*(17))"
	answer := Calculator(inputString)
	if answer != 2660 {
		t.Fatalf("Incorrect result")
	}
}

func TestCheckFail(t *testing.T) {
	var inputString = "19+6*8*6+18/2/3/3*9+4"
	answer := Calculator(inputString)
	if answer != 320 {
		t.Fatalf("fail")
	}
}
