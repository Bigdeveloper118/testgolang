package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	fmt.Print("Enter the number of rows: ")
	for y := 0; y < i; y++ {
		for z := 1; z <= y; z++ {
			fmt.Print(" ")
		}

		for x := i - y; x > 0; x-- {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
