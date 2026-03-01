// 8. Подсчёт чётных чисел
// Написать функцию countEvens(nums []int) int, 
// которая считает, сколько в слайсе чётных чисел. 

package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}

	fmt.Println(countEvens(s1))
}

func countEvens(nums []int)int {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] % 2 == 0 {
			count++
		}
	}

	return count
}