package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
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

type dollars float32
type databasess map[string]dollars

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

func (db databasess) ServerHttp(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func C7UseDb() {
	db := databasess{"shoes": 50, "socks": 5}
	http.HandleFunc("/", db.ServerHttp)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db databasess) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db databasess) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func C7UseDb2() {
	db := databasess{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/list", db.list)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
