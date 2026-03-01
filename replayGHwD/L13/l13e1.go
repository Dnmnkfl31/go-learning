// 1. Смены знака
// Подсчитать, сколько раз у соседних пар меняется знак.

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

	count := 0

	for i := 0; i < len(array) - 1; i++ {
		if array[i] > 0 && array[i + 1] < 0 || array[i] < 0 && array[i + 1] > 0 {
			count++
		}
	}

	fmt.Println(count)
}