package main

import (
	"fmt"
	"math"
	"os"

	"golang.org/x/net/html"
)

func C5Add(x int, y int) int {
	return x + y
}

func C5Sub(x int, y int) int {
	return x - y
}

func C5PrintFunc() {
	fmt.Printf("%T\n", C5Add)
	fmt.Printf("%T\n,", C5Sub)
	fmt.Printf("%T\n", math.Sin)
}

func C5Visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = C5Visit(links, c)
	}
	return links
}

func C5FindLink() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("doc: ", doc)
	fmt.Println()

	for _, link := range C5Visit(nil, doc) {
		fmt.Println(link)
	}
}

func C5ErrorF() {
	var s string = "ss"
	err := fmt.Errorf("getting : %s", s)
	fmt.Println(err)
}
func RawReturn() (a int, b int) {
	a = 1
	b = 2
	return
}

func C5UseRawReturn() {
	fmt.Println(RawReturn())
}
