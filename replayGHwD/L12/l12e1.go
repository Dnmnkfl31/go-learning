// 1. Замена нулей
// Заменить все элементы, равные нулю, на число 1 и вывести обновлённый массив.

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

	for i := 0; i < (len(array)); i++ {
		if array[i] == 0 {
			array[i] = 1
		}
	}

	fmt.Println("Новый массив:", array)
}