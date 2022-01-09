package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
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

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func C8ConcurrentServer() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// go handleConn(conn)
		handleConn(conn)
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func C8NetCat() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}

func C8UsePipeLine() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 30; x++ {
			naturals <- x
			time.Sleep(100 * time.Millisecond)
		}
		close(naturals)
	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()
	for {

		x, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(x)
	}
}
