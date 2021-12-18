package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func C4UseArray() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
	d := [2]int{1, 2}
	fmt.Println(a == d)

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

func C4P4_2() {
	flagSha := os.Args[2]

	if flagSha == "256" {
		fmt.Printf("%x\n", sha256.Sum256([]byte(os.Args[1])))
	} else if flagSha == "512" {
		fmt.Printf("%x\n", sha512.Sum512([]byte(os.Args[1])))
	} else if flagSha == "384" {
		fmt.Printf("%x\n", sha512.Sum384([]byte(os.Args[1])))
	}
}

func C4UseSlice() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

}

func C4AppendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func C4UseAppendInt() {
	a1 := []int{1, 2, 3}
	var y int = 10
	a2 := C4AppendInt(a1, y)
	fmt.Println(a1)
	fmt.Println(a2)
}

func C4ShowSliceCapChange() {
	var s []int
	for i := 0; i < 10; i++ {
		fmt.Println("cap： ", cap(s), "  len: ", len(s))
		s = append(s, i)
	}
}

func C4Remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func C4UseRemove() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(C4Remove(s, 2))
}

func C4PrintMap(m map[string]int) {
	for k, v := range m {
		fmt.Println("key: ", k, " , value: ", v)
	}
	fmt.Println(m)
	fmt.Println()
}

func C4UseMap() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
		"gg":      43,
		"tt":      71,
		"ll":      20,
	}

	// C4PrintMap(ages)
	// delete(ages, "s")
	// C4PrintMap(ages)
	// delete(ages, "alice")
	// C4PrintMap(ages)

	for i := 0; i < 10; i++ {
		C4PrintMap(ages)
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)

	for i := 0; i < 10; i++ {
		for _, name := range names {
			fmt.Printf("%s\t%d\n", name, ages[name])
		}
		fmt.Println()
	}

	names2 := make([]string, 0, len(ages))
	fmt.Println(len(names2), " ", cap(names2))
}

func C4UseJson() {
	type Movie struct {
		Title  string
		Year   int  `json:"released"`
		Color  bool `json:"color,omitempty"`
		Actors []string
	}

	var movies = []Movie{
		{Title: "a", Year: 1942, Color: true, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	}

	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

}
