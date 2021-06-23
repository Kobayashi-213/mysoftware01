package main

import "testing"

func Test01(t *testing.T) {
	int := solveString("X")
	if int != 10 {
		t.Error("Test01 is failed")
	}
}

func Test02(t *testing.T) {
	result := judgeSameStr("XXX", "XXX")
	expected := true
	if result != expected {
		t.Error("Test02 is failed")
	}
}
func Test03(t *testing.T) {
	result := solveString("MCMXC")
	expected := 1990
	if result != expected {
		t.Error("Test03 is failed")
	}
}
func Test04(t *testing.T) {
	result := solveString("XCIX")
	expected := 99
	if result != expected {
		t.Error("Test04 is failed")
	}
}

func Test05(t *testing.T) {
	result := solveString("XLVII")
	expected := 47
	if result != expected {
		t.Error("Test05 is failed")
	}
}
