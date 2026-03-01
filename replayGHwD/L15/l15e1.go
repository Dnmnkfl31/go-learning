// 1. Возведение числа в степень
// Написать функцию power(base int, exp int) int, 
// которая возводит число base в степень exp (предполагаем, что exp ≥ 0). 

package main

import "fmt"

func main() {
	fmt.Println(power(2, 3))
}

func power(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}

	return result
}