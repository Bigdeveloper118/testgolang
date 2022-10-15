package main

import (
	"fmt"
)

func main() {
	var i int
	var g string
	fmt.Scan(&i)
	fmt.Scan(&g)
	fmt.Print("Enter the number of rows: Enter print symbol: ")
	for y := 1; y <= i; y++ {
		for z := i - y; z > 0; z-- {
			fmt.Print(" ")
		}
		for t := 1; t <= (y*2)-1; t++ {
			fmt.Printf("%v", g)
		}
		fmt.Print("\n")
	}
}