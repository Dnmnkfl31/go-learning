// 14. Сумма положительных чисел
// Написать функцию sumPositive(nums []int) int, 
// которая возвращает сумму только положительных чисел из слайса.

package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(sumPositive(s1))
}

func sumPositive(nums []int)int {
	sumPozitiveNum := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] % 2 == 0 {
			sumPozitiveNum += nums[i]
		}
	}

	return sumPozitiveNum
}