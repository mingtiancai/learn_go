package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TestPrint() {
	fmt.Println("test")
}

func Echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)
}

func Echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func Echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//1.1
func P1_1() {
	fmt.Println(os.Args[0])
}

//1.2
func P1_2() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, " ", arg)
	}
}

func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
