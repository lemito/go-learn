package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

type Vertex struct {
	Lat, Lon float64
}

var b map[string]Vertex // key-string, val-vertex
var mm map[string]Vertex

func test() {
	b = make(map[string]Vertex)
	b["meow"] = Vertex{
		12.25, 67.89,
	}
	mm = map[string]Vertex{
		"woof": {122, 55},
	}
	fmt.Println(b)
	fmt.Println(mm)

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	arr := strings.Fields(s)
	for _, v := range arr {
		m[v]++;
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
