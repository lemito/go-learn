package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string){
	return y, x
}

func main() {
	fmt.Println("hello world")
	fmt.Println(time.Now())
	fmt.Printf("Rand float is = %f\n", rand.Float64())
	fmt.Printf("PI is = %f\n", math.Pi)
	fmt.Printf("%d + %d = %d\n", 5, 3, add(5, 3))
	a, b := swap("a", "b");
	var sum, i int64;
	for i = 0; i < 5; i++ {
		sum += i;
	}
	fmt.Println(sum);
	fmt.Println(a, " ", b);
}
