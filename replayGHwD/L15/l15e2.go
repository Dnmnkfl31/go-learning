// 2. Максимум из двух чисел
// Написать функцию max(a int, b int) int, которая возвращает большее из двух чисел.

package main

import "fmt"

func main() {
	fmt.Println(max(10, 2))
}

func max(a, b int)int {
	if a > b {
		return a

	} else {
		return b
	}
}