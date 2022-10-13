package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	fmt.Print("Enter the number of rows: ")
	for j := 1; j <= i; j++ {
		for k := 1; k <= j; k++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
