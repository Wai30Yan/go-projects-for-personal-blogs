package main

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	n := len(arr)
	mid := n / 2
	a1 := arr[:mid]
	a2 := arr[mid:]

	return merge(mergeSort(a1), mergeSort(a2))
}


func merge(a1, a2 []int) (result []int) {
	length := len(a1) + len(a2)
    result = make([]int, length)

    i := 0
    for len(a1) > 0 && len(a2) > 0 {
        if a1[0] < a2[0] {
            result[i] = a1[0]
            a1 = a1[1:]
        } else {
            result[i] = a2[0]
            a2 = a2[1:]
        }
        i++
    }

    for j := 0; j < len(a1); j++ {
        result[i] = a1[j]
        i++
    }
    for j := 0; j < len(a2); j++ {
        result[i] = a2[j]
        i++
    }

    return 
}
