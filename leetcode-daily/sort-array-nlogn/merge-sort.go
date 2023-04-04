package main

import "fmt"

func main() {
	slice := []int{4, 6, 3, 5, 8, 1, 2}
	fmt.Println(sortArray(slice))
}

func sortArray(nums []int) []int {
	tempArr := make([]int, len(nums))
	mergeSort(nums, 0, len(nums)-1, tempArr)
	return nums
}

func merge(arr []int, left, mid, right int, tempArr[]int) {
	start1, start2 := left, mid+1
	n1, n2 := mid-left+1, right-mid

	for i := 0; i < n1; i++ {
		tempArr[start1 + i] = arr[start1+i]
	}
	for i := 0; i < n2; i++ {
		tempArr[start2 + i] = arr[start2 + i]
	}

	i, j, k := 0, 0, left
	for i < n1 && j < n2 {
		if (tempArr[start1 + i] <= tempArr[start2 + j]) {
			arr[k] = tempArr[start1 + i]
			i += 1
		} else {
			arr[k] = tempArr[start2 + j]
			j += 1
		}
		k += 1
	}

	for i < n1 {
		arr[k] = tempArr[start1 + i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = tempArr[start2 + j]
		j++
		k++
	}
}

func mergeSort(arr []int, left, right int, tempArr []int) {
	if left >= right {
		return
	}

	mid := (left + right) / 2

	mergeSort(arr, left, mid, tempArr)
	mergeSort(arr, mid+1, right, tempArr)

	merge(arr, left, mid, right, tempArr)
}