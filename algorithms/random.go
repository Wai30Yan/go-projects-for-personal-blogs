package main

import (
	"math/rand"
	"time"
)

func random(min, max, size int) []int {
	arr := []int{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		num := rand.Intn(max-min) + min
		arr = append(arr, num)
	}
	return arr
}