// 3. Поиск элемента в слайсе
// Написать функцию contains(nums []int, target int) bool, 
// которая возвращает true, если элемент есть в слайсе, иначе false.

package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	fmt.Println(contains(s, 3))
	fmt.Println(contains(s, 5))
}

func contains(num []int, target int) bool {
	for i := 0; i < len(num); i++ {
		if num[i] == target {
			return true
		} 
	}
	return false
}