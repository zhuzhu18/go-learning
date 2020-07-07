package main

import "fmt"

func printArray(arr *[5]int) {
	arr[0] = 100
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func main() {
	var arr1 [3]int
	arr2 := [2]int{2, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [3][4]int

	fmt.Println(arr1, arr2, arr3, grid)
	printArray(&arr3)
}
