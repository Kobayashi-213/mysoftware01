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

func solve(n int) string {
	//千のくらいについてローマ数字を返す
	thouSandStr := getThousand(n)

	return thouSandStr
}

func main() {
	solve(1500)
}
