// 8. Проверка отсортированности
// Определить, является ли массив возрастающим, убывающим или неотсортированным.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	array := [10]int{}

	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(10) - 5
	}

	fmt.Println("Array:", array)

    ascending := true
    descending := true

    for i := 0; i < len(array)-1; i++ {
        if array[i] > array[i+1] {
            ascending = false
        }
        if array[i] < array[i+1] {
            descending = false
        }
    }

    if ascending {
        fmt.Println("Массив возрастающий")

    } else if descending {
        fmt.Println("Массив убывающий")
		
    } else {
        fmt.Println("Массив неотсортированный")
    }
}