package main

import "testing"

func Test01(t *testing.T) {
	str := solve(3000)
	if str != "MMM" {
		t.Error("Test01 is failed")
	}
}




/*func Test02(t *testing.T) {
	result := calcu(1990)
	expected := 1
	if result != expected {
		t.Error("Test02 is failed")
	}
}

func Test03(t *testing.T) {
	result := calcu(1990)
	result2 := numToStr(result)
	expected := "I"
	if result != expected {
		t.Error("Test03 is failed")
	}
}
func Test04(t *testing.T) {
	result := calcu(3000)
	result2 := numToStr(result)
	expected := "MMM"
	if result != expected {
		t.Error("Test03 is failed")
	}
}*/
