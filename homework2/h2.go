package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	fmt.Print("Enter the number of rows: ")
	for y := 1; y <= i; y++ {
		for z := i - y; z > 0; z-- {
			fmt.Print(" ")
		}
		for x := 1; x <= y; x++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
