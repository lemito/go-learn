package main

import "fmt"

func foo() {
	// паникуем
	panic("unexpected!")
}

func div(a, b int64) int64 {
	if (b == 0) {
		panic("Деление на нуль")
	}
	return a / b;
}

func main() {
	// выполняется при завершении main
	defer func() {
		// вызываем recover и сравниваем результат с nil
		if r := recover(); r != nil {
			fmt.Println(r) // выведет "unexpected"
		}
	}()
	div(10, 0)
	// foo() // внутри foo срабатывает паника
	fmt.Println("Вы не увидите это сообщение, так как в foo возникла паника")
}
