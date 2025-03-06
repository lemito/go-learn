package main

import (
	"fmt"
	"math"
)

func info_slice(a []int) {
	fmt.Printf("cap=%d len=%d %v \n", cap(a), len(a), a)
}

func main() {
	numbers := [6]int{1, 2, 3, 4, 5, 6}

	var num2 []int = numbers[2:4] // slice === указатель на массив
	var pNum []int = numbers[:]   // pNum - указатель, не копия
	fmt.Printf("%T\n", pNum)
	fmt.Println(num2)
	num2[0] = 9 // изменяем исходный
	fmt.Println(numbers)
	info_slice(pNum)
	info_slice(num2)

	a := make([]int, 0, 5) // create slice len=0 cap=5

	a = append(a, 1)
	a = append(a, 5)
	a = append(a, 7, 8, 9, 4, 5, 6, 1, 2, 3)

	info_slice(a)

	for _, v := range a {
		fmt.Printf("2**%d = %f\n", v, math.Pow(2, float64(v)))
	}
}
