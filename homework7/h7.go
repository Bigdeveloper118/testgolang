package main

import "fmt"

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func main() {
	var min int
	var max int

	size := 6

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("M[1][1]: M[1][2]: M[2][1]: M[2][2]: M[3][1]: M[3][2]:")
	fmt.Println("Matrix")

	fmt.Printf("%v\t%v\n%v\t%v\n%v\t%v\n\n", arr[0], arr[1], arr[2], arr[3], arr[4], arr[5])
	min, max = MinMax(arr)
	fmt.Printf("Min = %v\n", min)
	fmt.Printf("Max = %v\n", max)
}
