package main

import "fmt"

func main() {
	arr1 := random(1, 15, 10)
	fmt.Println(arr1)

	quickSort(arr1, 0, len(arr1)-1)
	fmt.Println(arr1)
}