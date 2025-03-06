package main

import (
	"fmt"
)

type word string

type Fruit string
type Name string

func main() {
	var str string
	str = "Hello, world!"

	var fruit Fruit
	var name Name
	fruit = "Banana"
	// name = fruit - нельзя, фрукт не имя
	name = "Qwerty"

	fmt.Printf("%s %s\n", fruit, name)

	println(word(str[0])) // H
	println(str[0])       // 72



}
