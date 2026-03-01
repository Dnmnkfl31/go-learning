// 9. Удвоение элементов
// Создать новый массив того же размера, где каждый элемент равен удвоенному значению исходного, и вывести его.

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

	newArray := [10]int{}

	for i := 0; i < len(newArray); i++ {
		newArray[i] = array[i] * 2
	}

	fmt.Println("Удвоенный массив:"newArray)
}