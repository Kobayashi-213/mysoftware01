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

func solve(n int) string {
	//千のくらいについてローマ数字を返す
	thouSandStr := getThousand(n)

	hnum := calcu(n, 1000)

	hundRedStr := getHundred(hnum)

	return thouSandStr + hundRedStr
}

func calcu(n int, m int) int {
	return (n % m)
}

func main() {
	solve(1500)
}
