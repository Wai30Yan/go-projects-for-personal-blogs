package main

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
		return false
	}

	mST, mTS := make(map[rune]rune), make(map[rune]rune)

	for i := 0; i < len(s); i++ {
		c1 := rune(s[i])
		c2 := rune(t[i])

		_, ok1 := mST[c1]
		_, ok2 := mTS[c2]

		if !ok1 && !ok2 {
			mST[c1] = c2
			mTS[c2] = c1
		} else if !(mST[c1] == c2 && mTS[c2] == c1) {
			return false
		} 
	}

	return true
}