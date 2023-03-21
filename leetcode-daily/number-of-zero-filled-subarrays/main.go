package main

func main() {

}

func zeroFilledSubarray(nums []int) int64 {
    var count, subcount int64 = 0, 0
    for _, num := range nums {
        if num != 0 {
            subcount = 0
            continue
        }

        subcount++
        count += subcount
    }
    return count
}