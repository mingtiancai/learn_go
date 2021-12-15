package main

import (
	"fmt"
	"math"
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
