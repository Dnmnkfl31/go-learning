// 5. Соседи одинакового знака
// Посчитать, сколько раз встречается пара соседних элементов, 
// у которых одинаковый знак (оба положительные или оба отрицательные).

package main

import(
	"fmt"
	"math/rand"
)

func main() {
	array := [10]int{}

	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(10) - 5
	}

	fmt.Println("Массив:", array)

	sum := 0

	for i := 0; i < len(array) - 1; i++ {
		if array[i] < 0 && array[i + 1] < 0 {
			sum++

		} else if array[i] > 0 && array[i + 1] > 0 {
			sum++
		}
	}

	fmt.Println(sum)
}