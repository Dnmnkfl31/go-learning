// 2. Сумма элементов
// Вычислить сумму всех элементов массива и вывести результат.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := [10]int{}

	for i := 0; i < 10; i++ {
		array[i] = rand.Intn(i + 1)
	}

	fmt.Println("Начальный массив:", array)

	sum := 0

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	fmt.Println(sum)
}