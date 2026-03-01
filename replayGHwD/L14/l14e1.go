// Задача 1. Пустой слайс и append
// Создать пустой слайс var s []int. С помощью append добавить в него числа 5, 10 и 15. 
// Вывести получившийся слайс вместе с длиной (len) и ёмкостью (cap).

package main

import "fmt"

func main() {
	var s []int

	s = append(s, 5, 10, 15)

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}