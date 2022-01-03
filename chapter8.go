package main

import (
	"fmt"
	"time"
)

func C8Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func C8Fib(x int) int {
	if x < 2 {
		return x
	}
	return C8Fib(x-1) + C8Fib(x-2)
}

func C8UseFib() {
	go C8Spinner(100 * time.Microsecond)
	const n = 45
	fibN := C8Fib(n)
	fmt.Printf("\rFibonacci(%d)=%d\n", n, fibN)
}
