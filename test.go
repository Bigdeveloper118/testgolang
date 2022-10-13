package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	fmt.Print("Enter the number of rows: ")
	for k := 1; k <= i; k++ {
		for j := 1; j <= k; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
