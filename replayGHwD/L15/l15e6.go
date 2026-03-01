// 6. Минимальный элемент слайса
// Написать функцию minSlice(nums []int) int, 
// которая возвращает минимальный элемент слайса. Если слайс пустой — вернуть 0. 

package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{}

	fmt.Println(minSlice(s1))
	fmt.Println(minSlice(s2))
}

func minSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minElement := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] < minElement {
			minElement = nums[i]
		}
	}

	return minElement
}