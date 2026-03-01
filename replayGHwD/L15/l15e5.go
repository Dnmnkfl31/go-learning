// 5. Среднее арифметическое
// Написать функцию average(nums []int) float64, 
// которая возвращает среднее арифметическое элементов слайса. 

package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	fmt.Println(average(s))
}

func average(nums []int) float64 {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	result := float64(sum) / float64(len(nums))

	return result
}