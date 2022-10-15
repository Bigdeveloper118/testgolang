package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	fmt.Print("Enter the number of rows: ")
	for z := 1; z <= i; z++ {
		for y := 1; y <= z; y++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
