package main


func main() {
	maxWeight, numOfItems := 8, 4
	profit := []int{0,1,2,5,6}
	weight := []int{0,2,3,4,5}

	knapsack(weight, profit, maxWeight, numOfItems)
}