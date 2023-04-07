package main

func minEatingSpeed(piles []int, h int) int {
    left, right := 1, 1000000000

	for left < right {
		mid := left + (right-left) / 2

		if canEatAll(piles, h, mid) {
			right = mid
		} else {
			left = mid+1
		}
	}

	return left
}

func canEatAll(piles []int, h, k int) bool {
	time := 0
	for _, p := range piles {
		time += (p-1)/k + 1
	}
	return time <= h
}