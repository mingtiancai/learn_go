package main

import (
	"flag"
	"fmt"
	"strings"
)

const boilingF = 121.0

func C2UseConst() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g F or %g C\n", f, c)

}

func C2F() *int {
	i := 1
	return &i
}

func C2UsePtr() {
	x := 1
	p := &x
	fmt.Println(*p)
	fmt.Println(p)
	*p = 2
	fmt.Println(x)

	//test local variant ptr
	fmt.Println("C2F: ", C2F(), "  ", *C2F())
	fmt.Println("C2F: ", C2F(), "  ", *C2F())
	fmt.Println("C2F: ", C2F(), "  ", *C2F())
}

func C2Echo4() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", "    ", "separator")
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func C2UseFlag() {
	host := flag.String("host", "127.0.0.1", "请输入host地址")
	port := flag.Int("port", 3306, "请输入端口号")
	flag.Parse() // 解析参数
	fmt.Printf("%s:%d\n", *host, *port)
}

type Celsius float64
type Fahrenheit float64

func C2UseType() {
	var c Celsius = 10.0
	var f Fahrenheit = 12.0
	fmt.Println("c: ", c)
	//fmt.Println(c - f)  //即使地层类型相同也不可以进行云运算
	fmt.Println("f: ", f)
	f = Fahrenheit(c)
	fmt.Println("f: ", f)

}
