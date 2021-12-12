package main

import (
	"fmt"
)

const boilingF = 121.0

func C2UseConst() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F or %g C\n", f, c)

}

func C2UsePtr() {
	x := 1
	p := &x
	fmt.Println(*p)
	fmt.Println(p)
	*p = 2
	fmt.Println(x)

}