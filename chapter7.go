package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func C7UseByteCounter() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	name := "do"
	fmt.Fprintf(&c, "hello,%s", name)
	fmt.Println(c)
}

func (r *C6Rocket) String() string {
	res := "name:  " + r.name + " length:  " + strconv.FormatFloat(r.length, 'E', -1, 64)
	return res
}

func C7UseRocket() {
	rocket := C6Rocket{"changzheng", 10}
	fmt.Println(rocket)
	fmt.Println(rocket.String())
}

const (
	n111 = 10
)

func C7UseFlag() {
	var C7period = flag.Duration("period", n111*time.Second, "sleep period")
	flag.Parse()
	fmt.Printf("sleeping for %v...", *C7period)
	time.Sleep(*C7period)
	fmt.Println()
}
