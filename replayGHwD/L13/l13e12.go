// 12. Соседние суммы равн
// Проверить, есть ли такие три подряд идущих элемента, 
// что сумма первых двух равна третьему.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := [10]int{}

	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(11) - 5
	}

	fmt.Println(array)

	for i := 0; i < len(array) - 2; i++ {
		if array[i] + array[i + 1] == array[i + 2] {
			fmt.Println("Есть")
			return
		}
	}

	fmt.Println("Такого случая нет")
}