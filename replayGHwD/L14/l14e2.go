// Задача 2: Добавление элементов и промежуточный вывод
// Создать пустой слайс var s []int. С помощью append добавить в него числа от 1 до 5.
// После каждого добавления выводить текущее состояние слайса.

package main

import "fmt"

func main() {
	var s []int

	s = append(s, 1)
	fmt.Println(s)

	s = append(s, 2)
	fmt.Println(s)

	s = append(s, 3)
	fmt.Println(s)

	s = append(s, 4)
	fmt.Println(s)

	s = append(s, 5)
	fmt.Println(s)
}