package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"time"

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

func C5WaitForServer(url string) error {
	const timeout = 5 * time.Second
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		resp, err := http.Get(url)
		if err == nil {
			fmt.Println(resp)
			return nil
		}
		log.Printf("server not responding (%s) ; retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func C5UseWait() {
	err := C5WaitForServer("https://ac.scmor.com/")
	if err != nil {
		fmt.Println("get failed")
	}

}
func C5Add1(r rune) rune {
	return r + 1
}

func C5Sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func C5Min(args ...int) (int, error) {
	n := len(args)
	fmt.Println("n: ", n)
	if n == 0 {
		err := fmt.Errorf("no args")
		return 0, err
	} else {
		var result int = args[0]
		for i := 0; i < n; i++ {
			if args[i] < result {
				result = args[i]
			}
		}
		return result, nil
	}
}

func C5UseFunc() {
	fmt.Println(strings.Map(C5Add1, "Hal"))
	x := []string{
		"ss",
		"saas",
	}

	fmt.Println(x)

	fmt.Println(C5Sum(1, 2, 3))
	fmt.Println(C5Sum(1, 2))
	fmt.Println(C5Min())
	fmt.Println(C5Min(1, 2, 3))

}

func C5Title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s,not text/html", url, ct)

	}
	// doc, err := html.Parse(resp.Body)
	// resp.Body.Close()

	// if err != nil {
	// 	return fmt.Errorf("parseing %s as HTML: %v", url, err)

	// }

	// visitNode := func(n *html.Node) {
	// 	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
	// 		fmt.Println(n.FirstChild.Data)
	// 	}
	// }
	// forEachNode(doc, visitNode, nil)
	return nil
}
