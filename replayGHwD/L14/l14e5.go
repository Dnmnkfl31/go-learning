// Задача 5. Увеличение capacity
// Создать слайс с помощью make([]int, 0, 2).
// В цикле добавить в него числа от 1 до 6, после каждого добавления выводя слайс, 
// его длину и ёмкость.
// Проследить, как Go автоматически увеличивает ёмкость при переполнении.

package main

import "fmt"

func main() {
	s := make([]int, 0, 2)

	for i := 1; i <= 6; i++ {
		s = append(s, i)
		fmt.Println("slice:", s)
		fmt.Println("cap =", cap(s))
		fmt.Println("len =", len(s))
	}
}