// 7. Поиск числа
// Запросить у пользователя число и вывести "Найдено", если оно есть в массиве, и "Не найдено" в противном случае.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := [10]int{}

	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(i + 1)
	}

	fmt.Println("Начальный массив:", array)

	fmt.Println("Введите n:")

	var n int
	fmt.Scanln(&n)

	for i := 0; i < len(array); i++ {
		if array[i] == n {
			fmt.Println("Найдено")
			return
		}
	}

	fmt.Println("Не найдено")
}