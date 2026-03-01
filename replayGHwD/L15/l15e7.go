// 7. Слияние двух слайсов
// Написать функцию merge(a []int, b []int) []int, которая объединяет два слайса в один.

package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}

	fmt.Println(merge(s1, s2))
}

func merge(s1 []int, s2 []int) []int {
	s3 := append(s1, s2...)
	return s3
}