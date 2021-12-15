package main

import "fmt"

func C3UseBit() {
	var x uint8 = 1<<1 | 1<<5
	fmt.Println(x)
	fmt.Printf("%08b\n", x)
}

func C3UseFor() {
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i])
	}
}

func C3UsePrint() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o)

	ascii := 'a'
	unicode := '国'
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)
}
