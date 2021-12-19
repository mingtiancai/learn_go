package main

import (
	"fmt"
	"math"
)

type C6Point struct{ x, y float64 }

func (p C6Point) Distance(q C6Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p *C6Point) C6ScaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

func C6UseMethod() {
	p := C6Point{1, 2}
	q := C6Point{1, 2}
	fmt.Println(p.Distance(q))

	pp := &p
	pp.C6ScaleBy(0.1)
	fmt.Println(*pp)
}

type Path []C6Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

func C6UsePath() {
	perim := Path{
		{1, 2},
		{2, 2},
		{3, 4},
	}
	fmt.Println(perim.Distance())
}
