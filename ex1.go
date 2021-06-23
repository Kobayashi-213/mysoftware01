package main

import (
	"fmt"
)

//千のくらいについてローマ数字を返す
func getThousand(n int) string {

	thouSand := (int)(n / 1000)
	fmt.Println(thouSand)
	str := ""
	//商の分によってローマ数字を
	if thouSand == 1 {
		str = "M"
	} else if thouSand == 2 {
		str = "MM"
	} else if thouSand == 3 {
		str = "MMM"
	} else if thouSand == 4 {
		str = "MMMM"
	}
	fmt.Println(str)
	return str
}

func getHundred(n int) string {
	hundred := (int)(n / 100)
	str := ""
	if hundred == 1 {
		str = "C"
	} else if hundred == 2 {
		str = "CC"
	} else if hundred == 3 {
		str = "CCC"
	} else if hundred == 4 {
		str = "CD"
	} else if hundred == 5 {
		str = "D"
	} else if hundred == 6 {
		str = "DC"
	} else if hundred == 7 {
		str = "DCC"
	} else if hundred == 8 {
		str = "DCCC"
	} else if hundred == 9 {
		str = "CM"
	}
	return str
}
func getTen(n int) string {
	ten := (int)(n / 10)
	str := ""
	if ten == 1 {
		str = "X"
	} else if ten == 2 {
		str = "XX"
	} else if ten == 3 {
		str = "XXX"
	} else if ten == 4 {
		str = "XL"
	} else if ten == 5 {
		str = "L"
	} else if ten == 6 {
		str = "LX"
	} else if ten == 7 {
		str = "LXX"
	} else if ten == 8 {
		str = "LXXX"
	} else if ten == 9 {
		str = "XC"
	}
	return str
}
func getOne(n int) string {
	one := (int)(n / 1)
	str := ""
	if one == 1 {
		str = "I"
	} else if one == 2 {
		str = "II"
	} else if one == 3 {
		str = "III"
	} else if one == 4 {
		str = "IV"
	} else if one == 5 {
		str = "V"
	} else if one == 6 {
		str = "VI"
	} else if one == 7 {
		str = "VII"
	} else if one == 8 {
		str = "VIII"
	} else if one == 9 {
		str = "IX"
	}
	return str
}

func solve(n int) string {
	//千のくらいについてローマ数字を返す
	thouSandStr := getThousand(n)

	hnum := calcu(n, 1000)

	hundRedStr := getHundred(hnum)
	tnum := calcu(n, 100)
	tenStr := getTen(tnum)

	onum := calcu(n, 10)
	oneStr := getOne(onum)
	return thouSandStr + hundRedStr + tenStr + oneStr
}

func calcu(n int, m int) int {
	return (n % m)
}

func main() {
	//fmt.Println(solve(1900))
}
