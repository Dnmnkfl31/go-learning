// 10. Подсчёт вхождений
// Запросить у пользователя число и подсчитать, сколько раз оно встречается в массиве.
// Вывести это количество.

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

	fmt.Println("Ввдеите n:")

	var n int
	fmt.Scanln(&n)

	count := 0

	for i := 0; i < len(array); i++ {
		if array[i] == n {
			count++
		}
	}

	fmt.Println(count)
}