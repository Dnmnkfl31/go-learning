// Задача 1. Проверка попадания в диапазон
// Условие: Пользователь вводит три числа: x, min, max. 
// Если x входит в диапазон включительно, то вывести "Попадает" Иначе — "Не попадает"

package main

import "fmt"

func main() {
	fmt.Println("Введите х:")

	var x int
	fmt.Scanln(&x)	

	fmt.Println("Введите min:")

	var min int
	fmt.Scanln(&min)	

	fmt.Println("Введите max:")

	var max int
	fmt.Scanln(&max)		

	if x > min && max > x {
		fmt.Println("Попадает")

	} else {
		fmt.Println("Не попадает")
	}
}