package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	fmt.Print("Enter number: The number is ")
	fmt.Print(i, "\n")
	if i%3 == 0 && i%5 == 0 {
		fmt.Print("FizzBuzz")
	} else if i%3 == 0 {
		fmt.Print("Fizz")
	} else if i%5 == 0 {
		fmt.Print("Buzz")
	}
}
