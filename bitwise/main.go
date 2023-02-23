package main

import (
	"fmt"
)

func main() {
	var x int8 = 1
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", x<<1)
	fmt.Printf("%08b\n", x<<5)
	fmt.Printf("%08b\n", x<<1 | x<<5)

	x = 1<<1 | 1<<5
	var y int8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	hex := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", hex)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

	ascii := 'a'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"

}

