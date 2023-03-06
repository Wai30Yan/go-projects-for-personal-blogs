package main

import "fmt"

func main() {
	arr1 := random(1, 15, 10)
	fmt.Println(arr1)

	arr2 := mergeSort(arr1)
	fmt.Println(arr2)
}