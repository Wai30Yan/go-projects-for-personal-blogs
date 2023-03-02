package main

import "fmt"

func main() {
	arr := random(1, 20, 10)
	fmt.Println(arr)
	arr2 := mergeSort(arr)
	fmt.Println(arr2)
}