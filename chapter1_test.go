package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	var a = 3
	var b = 4
	res := Sum(a, b)
	if res != 7 {
		t.Error("error")
	}
}
