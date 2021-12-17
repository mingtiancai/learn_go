package main

import "fmt"

func C4UseArray() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
	d := [2]int{1, 2}
	fmt.Println(a == d)
}
