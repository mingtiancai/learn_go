package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
	"time"
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

type C6ColoredPoint struct {
	C6Point
	Color color.RGBA
}

func C6UseC6ColoredPoint() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = C6ColoredPoint{C6Point{1, 1}, red}
	var q = C6ColoredPoint{C6Point{5, 4}, blue}
	fmt.Println(p.Distance(q.C6Point))
	(&p).C6ScaleBy(2)
	q.C6ScaleBy(2)
	fmt.Println(p.Distance(q.C6Point))
}

type Cache struct {
	mu      sync.Mutex
	mapping map[string]string
}

func UseC6Cache() {
	cache := Cache{mapping: make(map[string]string)}
	cache.mu.Lock()
	fmt.Println("begin sleep! ")
	time.Sleep(8 * time.Second)
	cache.mapping["ss"] = "asa"
	fmt.Println(cache.mapping["ss"])
	fmt.Println("end sleep! ")
	cache.mu.Unlock()
}

type C6Rocket struct {
	name   string
	length float64
}

func (r *C6Rocket) Launch() {
	fmt.Println("launch")
	fmt.Println(r.name)
	fmt.Println(r.length)
}

func C6UseMethodVar() {
	r := C6Rocket{"ChangZheng", 17.0}
	fmt.Println("Begin timing")
	time.AfterFunc(5*time.Second, r.Launch)
	time.Sleep(7 * time.Second)

}
