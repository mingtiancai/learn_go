package main

import "fmt"

func C5Add(x int, y int) int {
	return x + y
}

func C5Sub(x int, y int) int {
	return x - y
}

func C5PrintFunc() {
	fmt.Printf("%T\n", C5Add)
	fmt.Printf("%T\n,", C5Sub)
}
