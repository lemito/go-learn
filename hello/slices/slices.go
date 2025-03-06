package main

import "fmt"


func main()  {
	numbers := [6]int{1,2,3,4,5,6};
	
	var num3 = numbers[2:4];
	fmt.Println(num3);
}