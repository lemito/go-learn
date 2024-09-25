package main

import (
	"fmt"
	// "go/hello/contacts" должно экспортироваться
	"math"
)

// enum
const (
	Black = "black"
	Gray  = "gray"
	White = "white"
)

// или

const (
	Red    = iota // 0
	Yellow        // 1
	Green         //2
)

var (
	ver string = "v0.0.1"
	id  int    = 0
)

const (
	pi float32 = 3.1415
)
const (
	i         = 10
	f         = 1.5
	i64 int64 = 88
)

const (
	one = 2*iota + 1
	three
	five
	seven
	nine
	eleven
)

func main() {
	e := math.E

	fmt.Println("ver =", ver, "id =", id, "pi =", pi)
	fmt.Println(e)

	fmt.Println(Black, Gray, White)
	fmt.Println(Red, Yellow, Green)

	fmt.Println(one, three, five, seven, nine, eleven)

	// contacts.SetSupport("Служба поддержки")
	// fmt.Println(contacts.GetContact())
	// fmt.Println("Email:", contacts.Email)
}
