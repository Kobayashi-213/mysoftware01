package main

import "testing"

func Test01(t *testing.T) {
	str := solve(3000)
	if str != "MMM" {
		t.Error("Test01 is failed")
	}
}

func Test02(t *testing.T) {
	result := calcu(1990, 1000)
	expected := 990
	if result != expected {
		t.Error("Test02 is failed")
	}
}

func Test03(t *testing.T) {
	result := solve(1900)
	expected := "MCM"
	if result != expected {
		t.Error("Test03 is failed")
	}
}
func Test04(t *testing.T) {
	result := solve(1999)
	expected := "MCMXCIX"
	if result != expected {
		t.Error("Test04 is failed")
	}
}
