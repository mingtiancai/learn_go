package main

import (
	"bytes"
	"fmt"
	"math"
	"time"
	"unicode/utf8"
)

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

func C3F(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)
}

const (
	width, height = 600, 300
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func C3Corner(i, j int) (float64, float64) {

	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := C3F(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func C3Surface() {

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg'+ "+
		"style='stroke: grey;fill: white; stroke-width: 0.7'"+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := C3Corner(i+1, j)
			bx, by := C3Corner(i, j)
			cx, cy := C3Corner(i, j+1)
			dx, dy := C3Corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)

		}
	}
	fmt.Println("</svg>")
}

func C3UseBool() {
	s := ""
	var t bool
	t = s != "" && s[0] == 'x'
	fmt.Println(t)
}

func C3UseUtf8() {
	s := "hello, 世界"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}

func C3IntsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func C3PrintInts() {
	fmt.Println(C3IntsToString([]int{1, -2, 3}))
}

type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func C3UseConst() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
	fmt.Println("FlagUp ", FlagUp)
	fmt.Println("FlagBroadcast ", FlagBroadcast)
	fmt.Println("FlagLoopback ", FlagLoopback)
	fmt.Println("FlagPointToPoint ", FlagPointToPoint)
	fmt.Println("FlagMulticast ", FlagMulticast)
}
