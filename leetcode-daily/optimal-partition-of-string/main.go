package main

func partitionString(s string) int {
    freq := make([]int, 26)
    var ans int = 1
    for i := 0; i < len(s); i++ {
        var ch byte = byte(s[i] - 'a')
        if freq[ch] > 0 {
            ans++
            for j:= 0; j < 26; j++ {
                freq[j] = 0
            }
        }
        freq[ch]++
    }
    return ans
}

// func partitionString(s string) int {
//     lastSeen := make([]int, 26)
// 	count := 1
// 	substringStart := 0

// 	for i, r := range s {

// 		if lastSeen[byte(r - 'a')] >= substringStart {
// 			count++
// 			substringStart = i
// 		}
// 		lastSeen[byte(r - 'a')] = i
// 	}
// 	return count
// }