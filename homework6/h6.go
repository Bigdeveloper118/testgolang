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
	for y := 0; y < i; y++ {
		for z := 1; z <= y; z++ {
			fmt.Print(" ")
		}
		for t := ((i - y) * 2) - 1; t > 0; t-- {
			fmt.Printf("%v", g)
		}
		fmt.Print("\n")
	}
}
